#!/bin/bash

tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test)
tsukid query customstaking validator --addr=$(tsukid keys show -a validator --keyring-backend=test)
tsukid query multistaking pools
tsukid query multistaking undelegations
tsukid query multistaking outstanding-rewards $(tsukid keys show -a validator --keyring-backend=test)
tsukid query multistaking compound-info $(tsukid keys show -a validator --keyring-backend=test)

tsukid tx multistaking upsert-staking-pool tsukivaloper1h3f7w7ekjnpfcyjhktq06n4rl8yued9a0ap468 --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking delegate tsukivaloper1h3f7w7ekjnpfcyjhktq06n4rl8yued9a0ap468 1000000ukex --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking undelegate tsukivaloper1h3f7w7ekjnpfcyjhktq06n4rl8yued9a0ap468 10000ukex --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking claim-undelegation 1 --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking claim-rewards --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking set-compound-info true "" --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes --broadcast-mode=block

tsukid keys add delegator1 --keyring-backend=test --home=$HOME/.tsukid
tsukid keys add delegator2 --keyring-backend=test --home=$HOME/.tsukid
tsukid tx bank send validator $(tsukid keys show -a delegator1 --keyring-backend=test --home=$HOME/.tsukid) 100ubtc,10000ukex --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx bank send validator $(tsukid keys show -a delegator2 --keyring-backend=test --home=$HOME/.tsukid) 1000000ukex --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking delegate tsukivaloper1h3f7w7ekjnpfcyjhktq06n4rl8yued9a0ap468 10ubtc --from=delegator1 --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking delegate tsukivaloper1h3f7w7ekjnpfcyjhktq06n4rl8yued9a0ap468 100ukex --from=delegator2 --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block

tsukid tx multistaking set-compound-info true "" --from=delegator1 --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes --broadcast-mode=block
tsukid tx bank send validator $(tsukid keys show -a delegator1 --keyring-backend=test --home=$HOME/.tsukid) 1000000v1/ukex --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking register-delegator --from=delegator1 --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block

tsukid query bank balances $(tsukid keys show -a delegator1 --keyring-backend=test)
tsukid query bank balances $(tsukid keys show -a delegator2 --keyring-backend=test)
tsukid query multistaking outstanding-rewards $(tsukid keys show -a delegator1 --keyring-backend=test)
tsukid query multistaking outstanding-rewards $(tsukid keys show -a delegator2 --keyring-backend=test)
tsukid query multistaking staking-pool-delegators tsukivaloper1h3f7w7ekjnpfcyjhktq06n4rl8yued9a0ap468