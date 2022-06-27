#!/bin/bash

tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test)
tsukid query customstaking validator --addr=$(tsukid keys show -a validator --keyring-backend=test)
tsukid query multistaking pools
tsukid query multistaking undelegations
tsukid query multistaking outstanding-rewards $(tsukid keys show -a validator --keyring-backend=test)
tsukid query multistaking compound-info $(tsukid keys show -a validator --keyring-backend=test)

tsukid tx multistaking upsert-staking-pool tsukivaloper1nwcljqs98zkr39pwenngquryaaueztv6ejtljt --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking delegate tsukivaloper1nwcljqs98zkr39pwenngquryaaueztv6ejtljt 1000000ukex --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking undelegate tsukivaloper1nwcljqs98zkr39pwenngquryaaueztv6ejtljt 10000ukex --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking claim-undelegation 1 --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking claim-rewards --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking set-compound-info true "" --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes --broadcast-mode=block