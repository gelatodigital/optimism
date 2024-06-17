// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Deploy.s.sol";
import "../lib/safe-contracts/contracts/libraries/MultiSend.sol";

contract UpdateRoles is Deploy {

    function updateGuardian() public {
        console.log("Starting guardian update");

        address payable storageSetter = getAddress("StorageSetter");
        if (storageSetter == address(0)) {
            storageSetter = payable(address(deployStorageSetter()));
            save("StorageSetter", storageSetter);
        }

        address multiSendAddress = getAddress("MultiSend");
        if (multiSendAddress == address(0)) {
            deployMultiSend();
        }

        _updateGuardian();

        console.log("Finished guardian update");
    }

    function _updateGuardian() internal broadcast {
        address proxyAdmin = mustGetAddress("ProxyAdmin");
        address payable storageSetter = mustGetAddress("StorageSetter");
        address multiSend = mustGetAddress("MultiSend");
        address payable superchainConfigProxy = mustGetAddress("SuperchainConfigProxy");
        address payable superchainConfigImplementation = mustGetAddress("SuperchainConfig");

        // Get current guardian
        SuperchainConfig superchainConfig = SuperchainConfig(superchainConfigProxy);
        address currentGuardian = superchainConfig.guardian();
        console.log("Current guardian: %s", currentGuardian);

        // Get new guardian
        address newGuardian = cfg.superchainConfigGuardian();
        console.log("New guardian to be set: %s", newGuardian);

        // Skip if the new guardian is the same as the current guardian
        if (currentGuardian == newGuardian) {
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
            _innerCallData: abi.encodeCall(SuperchainConfig.initialize, (newGuardian, false))
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
