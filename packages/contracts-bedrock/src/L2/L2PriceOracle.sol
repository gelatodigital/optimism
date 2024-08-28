// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { ISemver } from "src/universal/ISemver.sol";
import { Constants } from "src/libraries/Constants.sol";
import { PythStructs } from "@pythnetwork/pyth-sdk-solidity/PythStructs.sol";
import { PythErrors } from "@pythnetwork/pyth-sdk-solidity/PythErrors.sol";

/// @custom:proxied
/// @custom:predeploy 0x4200000000000000000000000000000000000420
/// @title L2PriceOracle
contract L2PriceOracle is ISemver {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    mapping(bytes32 => PythStructs.Price) private prices;

    /// @notice Address of the special depositor account.
    function DEPOSITOR_ACCOUNT() public pure returns (address addr_) {
        addr_ = Constants.DEPOSITOR_ACCOUNT;
    }

    function updatePrices(bytes32[] calldata _ids, PythStructs.Price[] calldata _prices) external {
        require(
            msg.sender == DEPOSITOR_ACCOUNT(),
            "L2PriceOracle: only the depositor account can update oracle prices"
        );
        require(
            _ids.length == _prices.length,
            "L2PriceOracle: array length mismatch"
        );
        for (uint256 i = 0; i < _ids.length; i++)
            prices[_ids[i]] = _prices[i];
    }

    function getPriceUnsafe(bytes32 id) public view returns (PythStructs.Price memory price_) {
        price_ = prices[id];
        if (price_.publishTime == 0)
            revert PythErrors.PriceFeedNotFound();
    }

    function getPriceNoOlderThan(bytes32 id, uint256 age) external view returns (PythStructs.Price memory price_) {
        price_ = getPriceUnsafe(id);
        if (price_.publishTime + age > block.timestamp)
            revert PythErrors.StalePrice();
    }
}
