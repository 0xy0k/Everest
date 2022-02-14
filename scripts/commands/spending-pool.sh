#!/bin/bash

tsukid tx spending create-spending-pool --name="validator-rewards-pool" --claim-start=$(($(date -u +%s))) --claim-end=$(($(date -u +%s) + 200)) --expire=$(($(date -u +%s) + 200)) --token="ukex" --rate=0.1 --vote-quorum="33" --vote-period="60" --vote-enactment="30" --owner-roles="" --owner-accounts=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --beneficiary-roles="2" --beneficiary-accounts="" --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes  --broadcast-mode=block 

tsukid tx spending deposit-spending-pool --name="validator-rewards-pool" --amount=1000000ukex --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block 

tsukid query spending pool-by-name validator-rewards-pool --home=$HOME/.tsukid
tsukid query spending pool-names
