#!/bin/bash

# TODO: add dynamic duration proposal set examples here


# create proposal for setting poor network msgs
tsukid tx customgov proposal set-proposal-duration-proposal UpsertDataRegistryProposal 300 --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=1000ukex --yes

# query for proposals
tsukid query customgov proposals

# vote on the proposal
tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 
# check votes
tsukid query customgov votes 1 
# wait until vote end time finish
tsukid query customgov proposals

# query proposal duration
tsukid query customgov proposal-duration UpsertDataRegistryProposal

# query all proposal durations
tsukid query customgov all-proposal-durations

# batch operation
tsukid tx customgov proposal set-batch-proposal-durations-proposal UpsertDataRegistry,SetNetworkProperty 300,300 --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=1000ukex --yes