# Update immutables in OP contracts

Updates are performed in alignment with suggested Optimism procedures: https://github.com/ethereum-optimism/superchain-ops/tree/main/tasks/eth/010-1-guardian-upgrade

General update process:

1. Upgrade proxy to the [StorageSetter](https://github.com/gelatodigital/raas-optimism/blob/fe7c8362cb7a1e9b3bbc99af18e48c17a4852ad8/packages/contracts-bedrock/src/universal/StorageSetter.sol#L11) implementation and call [setBytes32(0, 0)](https://github.com/gelatodigital/raas-optimism/blob/fe7c8362cb7a1e9b3bbc99af18e48c17a4852ad8/packages/contracts-bedrock/src/universal/StorageSetter.sol#L24) to reset the `initialized` slot.
2. Upgrade proxy back to the regular implementation  and re-initialize with new values.

Methods 1. and 2. are called via delegatecall from SystemOwnerSafe `execTransaction` to [MultiSend](https://github.com/safe-global/safe-contracts/blob/6b3875c99e234c5c26fa8fc1f712d054b2d45bcf/contracts/libraries/MultiSend.sol#L11) `multiSend` which in turns calls to [ProxyAdmin](https://github.com/gelatodigital/raas-optimism/blob/1d6894f43f55bc4e90dd497b7b471948df7706c0/packages/contracts-bedrock/src/universal/ProxyAdmin.sol#L30) `upgradeAndCall` method. The inner calldata of `upgradeAndCall` actually performs 1. and 2.

SuperchainConfig guardian update will update and reinitialize the SuperchainConfigProxy contract
L2OutputOracle challenger and submission interval update will update and reinitialize the L2OutputOracleProxy contract

## Prerequisites

1. Set up environment variables
    ```
    cp .env.example .env
    ```
    Set values:
   - ADMIN_PK: Private key of the admin account
   - RPC_URL: L1 RPC URL
   - DEPLOYMENT_CONTEXT: Deployment context of the chain
   - IMPL_SALT: Unique salt for the implementation deployment
   - ETHERSCAN_API_KEY: Etherscan API key for contract verification
2. Copy deploy config json to the `deploy-config` folder and use the same file name as in `DEPLOYMENT_CONTEXT`
3. Copy file with deployed contract addresses to the `deployments/$DEPLOYMENT_CONTEXT` folder and make sure to rename the file to `.deploy`
4. Update deploy config with the desired config changes. You can set the following variables:
   - `superchainConfigGuardian`
   - `l2OutputOracleChallenger`
   - `l2OutputOracleSubmissionInterval`

## Execute update for all values
This will apply changes to the contracts which deviate from the provided values in deploy config
```
./scripts/update-immutables.sh
```

