#!/bin/bash

# create proposal for setting poor network msgs
tsukid tx customgov proposal set-poor-network-msgs AAA,BBB --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=1000ukex --yes
# query for proposals
tsukid query customgov proposals
# set permission to vote on proposal
tsukid tx customgov permission whitelist --permission=19 --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes
# vote on the proposal
tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 
# check votes
tsukid query customgov votes 1 
# wait until vote end time finish
tsukid query customgov proposals
# query poor network messages
tsukid query customgov poor-network-messages

# whitelist permission for modifying network properties
tsukid tx customgov permission whitelist --from validator --keyring-backend=test --permission=7 --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
# test poor network messages after modifying min_validators section
tsukid tx customgov set-network-properties --from validator --min_validators="2" --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
# set permission for upsert token rate
tsukid tx customgov permission whitelist --from validator --keyring-backend=test --permission=$PermUpsertTokenRate --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
# try running upsert token rate which is not allowed on poor network
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="mykex" --rate="1.5" --fee_payments=true --chain-id=testing --fees=100ukex --home=$HOME/.tsukid  --yes
# try sending more than allowed amount via bank send
tsukid tx bank send validator $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) 100000000ukex --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
# try setting network property by governance to allow more amount sending
tsukid tx customgov proposal set-network-property POOR_NETWORK_MAX_BANK_SEND 100000000  --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes
tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes
# try sending after modification of poor network bank send param
tsukid tx bank send validator $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) 100000000ukex --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes