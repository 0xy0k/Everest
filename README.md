# Everest

Cross-chain Lending Aggregator

## Setup

1. Get required keys and set them as env variables

```
cp sample.env .env
source .env
```

2. Install Foundry
   Find the instructions [here](https://book.getfoundry.sh/getting-started/installation).

3. Install npm packages

```
yarn
```

4. Compile contracts

```
yarn build
```

5. Test

```
yarn test
```

## Git Submodules used in Foundry

This package works with git submodules (due to Foundry)

From time to time you may require to refresh/update your submodudles.
Follow this steps:

1. Clean the `lib` folder where foundry keeps git submodules, from the project root directory:

```
sudo rm -r ~/packages/protocol/lib/*
```

2. Run the command to update the submodule files:

```
git submodules update --recursive
```

## Slither

Slither is a Solidity static analyzer. It helps identify potential security issues.

1. [How to install](https://github.com/crytic/slither#how-to-install)

2. Run from "packages/protocol"

```
slither .
```

3. Run and generate outputs in a file

```
slither . --json slither_output.json
```

## Factors in contracts

A factor through these contracts refer to a fixed-digit decimal number. Specifically, a decimal number scaled by 1e18. These numbers should be treated as real numbers scaled down by 1e18. For example, the number 50% would be represented as 5\*1e17.

## Testing deployment of smart contracts

1. Ensure you have set `TEST_DEPLOY_NETWORK` in your .env file.

2. Compile the contracts:

```
forge build
```

3. Run a fork using the RPC of the same NETWORK as set in `TEST_DEPLOY_NETWORK`. For example, if your `TEST_DEPLOY_NETWORK`=goerli, then create your fork with:

```
anvil --chain-id 420 --fork-url https://goerli.infura.io/v3/<YOUR_KEY>
```

4. Deploy the system with the following command:

```
npx hardhat --network localhost deploy --no-compile
```

NOTE: The configuration in `hardhat.config.ts` uses the artifacts and bytecode created by the previously called `forge build` command.

5. A folder will be created under `deployments/localhost` with all the artifacts containing also the addresses.
