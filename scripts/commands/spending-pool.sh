#!/bin/bash

tsukid tx spending create-spending-pool --name="validator-rewards-pool" --claim-start=$(($(date -u +%s))) --claim-end=$(($(date -u +%s) + 200)) --expire=$(($(date -u +%s) + 200)) --token="ukex" --rate=0.1 --vote-quorum="33" --vote-period="60" --vote-enactment="30" --owner-roles="" --owner-accounts=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --beneficiary-roles="2" --beneficiary-accounts="" --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes  --broadcast-mode=block 

tsukid tx spending deposit-spending-pool --name="validator-rewards-pool" --amount=1000000ukex --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block 

tsukid tx spending register-spending-pool-beneficiary --name="validator-rewards-pool" --beneficiary-roles="1" --beneficiary-accounts="" --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block 
 
tsukid tx spending claim-spending-pool --name="validator-rewards-pool" --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block 

tsukid query spending pool-by-name validator-rewards-pool --home=$HOME/.tsukid
tsukid query spending pool-names

tsukid query customgov roles $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# proposals
tsukid tx spending proposal-update-spending-pool --title="title" --description="description" --name="validator-rewards-pool" --claim-start=$(($(date -u +%s))) --claim-end=$(($(date -u +%s) + 200)) --expire=$(($(date -u +%s) + 200)) --token="ukex" --rate=0.5 --vote-quorum="33" --vote-period="60" --vote-enactment="30" --owner-roles="" --owner-accounts=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --beneficiary-roles="2" --beneficiary-accounts="" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes
tsukid tx spending proposal-spending-pool-withdraw --title="title" --description="description" --name="validator-rewards-pool" --beneficiary-accounts=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --amount=210000ukex --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes
tsukid tx spending proposal-spending-pool-distribution --title="title" --description="description" --name="validator-rewards-pool" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes

tsukid query customgov proposals
tsukid query customgov proposal 1

tsukid tx customgov proposal vote 5 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 
