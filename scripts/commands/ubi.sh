#!/bin/bash

tsukid query customgov network-properties
tsukid query ubi ubi-record-by-name ValidatorQuaterlyIncome --home=$HOME/.tsukid
tsukid query ubi ubi-records

tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# set permissions
PermCreateUpsertUBIProposal=53
PermVoteUpsertUBIProposal=54
PermCreateRemoveUBIProposal=55
PermVoteRemoveUBIProposal=56
tsukid tx customgov proposal account whitelist-permission $PermCreateUpsertUBIProposal  --title="title" --description="description" --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block 
tsukid tx customgov proposal account whitelist-permission $PermVoteUpsertUBIProposal  --title="title" --description="description" --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block 
tsukid tx customgov proposal account whitelist-permission $PermCreateRemoveUBIProposal  --title="title" --description="description" --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block 
tsukid tx customgov proposal account whitelist-permission $PermVoteRemoveUBIProposal  --title="title" --description="description" --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block 

# proposals
tsukid tx customgov proposal set-network-property UBI_HARDCAP 20000000  --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block 
tsukid tx ubi proposal-upsert-ubi --title="title" --description="description" --name="ValidatorQuaterlyIncome" --distr-start=0 --distr-end=0 --amount=2000000 --period=7776000 --pool-name="ValidatorBasicRewardsPool" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block 
tsukid tx ubi proposal-remove-ubi --title="title" --description="description" --name="ValidatorQuaterlyIncome" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block 

tsukid query customgov proposals
tsukid query customgov proposal 1

tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes  --broadcast-mode=block 
tsukid tx customgov proposal vote 2 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes  --broadcast-mode=block 
tsukid tx customgov proposal vote 3 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes  --broadcast-mode=block 
