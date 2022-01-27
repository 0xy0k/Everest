#!/bin/bash

RoleValidator=2

tsukid tx customgov proposal account assign-role $RoleValidator  --title="title" --description="description" --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes

tsukid query customgov proposals
tsukid query customgov proposal 1

tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 

# check roles
tsukid query customgov roles $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)
