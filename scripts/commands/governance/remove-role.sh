#!/bin/bash

# for complexity of the operation, this proposal is not implemented
tsukid tx customgov proposal role remove-role newrole --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes

tsukid query customgov proposals
tsukid query customgov proposal 1

tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 

# check deleted role
tsukid query customgov role 3
tsukid query customgov role newrole
