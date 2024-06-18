// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Deploy.s.sol";
import "../lib/safe-contracts/contracts/libraries/MultiSend.sol";

contract UpdateImmutables is Deploy {

    function updateAll() public {
        updateL2OutputOracle();
        updateGuardian();
    }

    function updateChallenger() public {
        console.log("Starting L2OutputOracle challenger update");

        address payable storageSetter = getAddress("StorageSetter");
        if (storageSetter == address(0)) {
            storageSetter = payable(address(deployStorageSetter()));
            save("StorageSetter", storageSetter);
        }

        address multiSendAddress = getAddress("MultiSend");
        if (multiSendAddress == address(0)) {
            deployMultiSend();
        }

        // Get new challenger
        address newChallenger = cfg.l2OutputOracleChallenger();
        console.log("New challenger to be set: %s", newChallenger);

        L2OutputOracle l2OutputOracle = L2OutputOracle(mustGetAddress("L2OutputOracleProxy"));
        _updateL2OutputOracle(newChallenger, l2OutputOracle.submissionInterval());

        console.log("Finished L2OutputOracle challenger update");
    }

    function updateSubmissionInterval() public {
        console.log("Starting L2OutputOracle submission interval update");

        address payable storageSetter = getAddress("StorageSetter");
        if (storageSetter == address(0)) {
            storageSetter = payable(address(deployStorageSetter()));
            save("StorageSetter", storageSetter);
        }

        address multiSendAddress = getAddress("MultiSend");
        if (multiSendAddress == address(0)) {
            deployMultiSend();
        }

        // Get new submission interval
        uint256 newSubmissionInterval = cfg.l2OutputOracleSubmissionInterval();
        console.log("New submission interval to be set: %s", newSubmissionInterval);

        L2OutputOracle l2OutputOracle = L2OutputOracle(mustGetAddress("L2OutputOracleProxy"));
        _updateL2OutputOracle(l2OutputOracle.challenger(), newSubmissionInterval);

        console.log("Finished L2OutputOracle submission interval update");
    }

    function updateL2OutputOracle() public {
        console.log("Starting L2OutputOracle update");

        address payable storageSetter = getAddress("StorageSetter");
        if (storageSetter == address(0)) {
            storageSetter = payable(address(deployStorageSetter()));
            save("StorageSetter", storageSetter);
        }

        address multiSendAddress = getAddress("MultiSend");
        if (multiSendAddress == address(0)) {
            deployMultiSend();
        }

        // Get new challenger
        address newChallenger = cfg.l2OutputOracleChallenger();
        console.log("New challenger to be set: %s", newChallenger);

        // Get new submission interval
        uint256 newSubmissionInterval = cfg.l2OutputOracleSubmissionInterval();
        console.log("New submission interval to be set: %s", newSubmissionInterval);

        _updateL2OutputOracle(newChallenger, newSubmissionInterval);

        console.log("Finished L2OutputOracle update");
    }

    function updateGuardian() public {
        console.log("Starting SuperchainConfig guardian update");

        address payable storageSetter = getAddress("StorageSetter");
        if (storageSetter == address(0)) {
            storageSetter = payable(address(deployStorageSetter()));
            save("StorageSetter", storageSetter);
        }

        address multiSendAddress = getAddress("MultiSend");
        if (multiSendAddress == address(0)) {
            deployMultiSend();
        }

        // Get new guardian
        address newGuardian = cfg.superchainConfigGuardian();
        console.log("New guardian to be set: %s", newGuardian);

        _updateGuardian(newGuardian);

        console.log("Finished SuperchainConfig guardian update");
    }

    function _updateL2OutputOracle(address _newChallenger, uint256 _newSubmissionInterval) internal broadcast {
        address proxyAdmin = mustGetAddress("ProxyAdmin");
        address payable l2OutputOracleProxy = mustGetAddress("L2OutputOracleProxy");

        L2OutputOracle l2OutputOracle = L2OutputOracle(l2OutputOracleProxy);

        {
            // Get current values
            address currentChallenger = l2OutputOracle.challenger();
            console.log("Current challenger: %s", currentChallenger);
            uint256 currentSubmissionInterval = l2OutputOracle.submissionInterval();
            console.log("Current submission interval: %s", currentSubmissionInterval);

            // Skip updates if values are the same
            if (currentChallenger == _newChallenger && currentSubmissionInterval == _newSubmissionInterval) {
                console.log("New values are the same as the current values, skipping");
                return;
            }
        }

        // Build transaction to upgrade L2OutputOracleProxy to StorageSetter and call setBytes32(0, 0) to reset the "initialized" slot
        bytes memory setStorageData = _buildUpgradeAndCallData({
            _proxy: l2OutputOracleProxy,
            _implementation: mustGetAddress("StorageSetter"),
            _innerCallData: abi.encodeWithSignature(
                "setBytes32(bytes32,bytes32)",
                bytes32(0), bytes32(0)
            )
        });

        // Build transaction to upgrade L2OutputOracleProxy to the old L2OutputOracle implementation and initialize
        bytes memory initializeData = _buildUpgradeAndCallData({
            _proxy: l2OutputOracleProxy,
            _implementation: mustGetAddress("L2OutputOracle"),
            _innerCallData: abi.encodeCall(
                L2OutputOracle.initialize,
                (
                    _newSubmissionInterval,
                    l2OutputOracle.l2BlockTime(),
                    l2OutputOracle.startingBlockNumber(),
                    l2OutputOracle.startingTimestamp(),
                    l2OutputOracle.proposer(),
                    _newChallenger,
                    l2OutputOracle.finalizationPeriodSeconds()
                )
            )
        });

        // Build MultiSend batched transaction
        bytes memory multiSendData = bytes.concat(
            _buildMultiSendData({
                _target: proxyAdmin,
                _data: setStorageData
            }),
            _buildMultiSendData({
                _target: proxyAdmin,
                _data: initializeData
            }));

        // Execute MultiSend
        _delegateCallViaSafe({
            _target: mustGetAddress("MultiSend"),
            _data: abi.encodeCall(MultiSend.multiSend, (multiSendData))
        });

        console.log("Running chain assertions on the L2OutputOracle");
        console.log("New L2OutputOracle challenger: %s", l2OutputOracle.challenger());
        console.log("New L2OutputOracle submission interval: %s", l2OutputOracle.submissionInterval());
        require(loadInitializedSlot("L2OutputOracleProxy") == 1, "L2OutputOracleProxy is not initialized");
        require(l2OutputOracle.challenger() == cfg.l2OutputOracleChallenger());
        require(l2OutputOracle.submissionInterval() == cfg.l2OutputOracleSubmissionInterval());
    }

    function _updateGuardian(address _newGuardian) internal broadcast {
        address proxyAdmin = mustGetAddress("ProxyAdmin");
        address payable storageSetter = mustGetAddress("StorageSetter");
        address multiSend = mustGetAddress("MultiSend");
        address payable superchainConfigProxy = mustGetAddress("SuperchainConfigProxy");
        address payable superchainConfigImplementation = mustGetAddress("SuperchainConfig");

        // Get current guardian
        SuperchainConfig superchainConfig = SuperchainConfig(superchainConfigProxy);
        address currentGuardian = superchainConfig.guardian();
        console.log("Current guardian: %s", currentGuardian);

        // Skip if the new guardian is the same as the current guardian
        if (currentGuardian == _newGuardian) {
            console.log("New guardian is the same as the current guardian, skipping");
            return;
        }

        // Build transaction to upgrade SuperchainConfigProxy to StorageSetter and call setBytes32(0, 0) to reset the "initialized" slot
        bytes memory setStorageData = _buildUpgradeAndCallData({
            _proxy: superchainConfigProxy,
            _implementation: storageSetter,
            _innerCallData: abi.encodeWithSignature(
                "setBytes32(bytes32,bytes32)",
                bytes32(0), bytes32(0)
            )
        });

        // Build transaction to upgrade SuperchainConfigProxy to the old SuperchainConfig implementation and initialize
        bytes memory initializeData = _buildUpgradeAndCallData({
            _proxy: superchainConfigProxy,
            _implementation: superchainConfigImplementation,
            _innerCallData: abi.encodeCall(
                SuperchainConfig.initialize,
                (
                    _newGuardian,
                    false
                )
            )
        });

        // Build MultiSend batched transaction
        bytes memory multiSendData = bytes.concat(
            _buildMultiSendData({
                _target: proxyAdmin,
                _data: setStorageData
            }),
            _buildMultiSendData({
                _target: proxyAdmin,
                _data: initializeData
            }));

        // Execute MultiSend
        _delegateCallViaSafe({
            _target: multiSend,
            _data: abi.encodeCall(MultiSend.multiSend, (multiSendData))
        });

        console.log("Running chain assertions on the SuperchainConfig");
        console.log("New SuperchainConfig guardian: %s", superchainConfig.guardian());
        require(loadInitializedSlot("SuperchainConfigProxy") == 1, "SuperchainConfigProxy is not initialized");
        require(superchainConfig.guardian() == cfg.superchainConfigGuardian());
        require(superchainConfig.paused() == false);
    }

    /// @notice Call from the Safe contract to the Proxy Admin's upgrade and call method
    function _buildUpgradeAndCallData(address payable _proxy, address _implementation, bytes memory _innerCallData) internal pure returns (bytes memory) {
        return abi.encodeCall(ProxyAdmin.upgradeAndCall, (payable(_proxy), _implementation, _innerCallData));
    }

    function _buildMultiSendData(address _target, bytes memory _data) internal pure returns (bytes memory) {
        return abi.encodePacked(uint8(0), address(_target), uint256(0), uint256(_data.length), bytes(_data));
    }

    /// @notice Make a call from the Safe contract to an arbitrary address with arbitrary data
    function _multiSendViaSafe(address _target, bytes memory _data) internal {
        bytes memory data = abi.encodeCall(MultiSend.multiSend, (_data));
        _delegateCallViaSafe({_target: _target, _data: data});
    }

    /// @notice Make a call from the Safe contract to an arbitrary address with arbitrary data
    function _delegateCallViaSafe(address _target, bytes memory _data) internal {
        Safe safe = Safe(mustGetAddress("SystemOwnerSafe"));

        // This is the signature format used the caller is also the signer.
        bytes memory signature = abi.encodePacked(uint256(uint160(msg.sender)), bytes32(0), uint8(1));

        safe.execTransaction({
            to: _target,
            value: 0,
            data: _data,
            operation: SafeOps.Operation.DelegateCall,
            safeTxGas: 0,
            baseGas: 0,
            gasPrice: 0,
            gasToken: address(0),
            refundReceiver: payable(address(0)),
            signatures: signature
        });
    }

    function deployMultiSend() public broadcast returns (address addr_) {
        MultiSend multiSend = new MultiSend();

        console.log("MultiSend deployed at %s", address(multiSend));

        save("MultiSend", address(multiSend));
        addr_ = address(multiSend);
    }
}
