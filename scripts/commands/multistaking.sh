#!/bin/bash

tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test)
tsukid query customstaking validator --addr=$(tsukid keys show -a validator --keyring-backend=test)
tsukid query multistaking pools
tsukid query multistaking undelegations
tsukid query multistaking outstanding-rewards $(tsukid keys show -a validator --keyring-backend=test)

tsukid tx multistaking upsert-staking-pool tsukivaloper14sacxjzc936nhz27lhdq7gt0kaavxf886vr543 --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking delegate tsukivaloper14sacxjzc936nhz27lhdq7gt0kaavxf886vr543 1000000ukex --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking undelegate tsukivaloper14sacxjzc936nhz27lhdq7gt0kaavxf886vr543 10000ukex --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking claim-undelegation 1 --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking claim-rewards --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
