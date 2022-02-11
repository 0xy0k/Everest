#!/bin/bash

PermCreateUpsertDataRegistryProposal=10

tsukid tx customgov proposal role remove-whitelisted-permission newrole $PermCreateUpsertDataRegistryProposal  --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes

tsukid query customgov proposals
tsukid query customgov proposal 1

tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 

# check permissions
tsukid query customgov role newrole
