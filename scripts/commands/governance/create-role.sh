#!/bin/bash

tsukid tx customgov proposal role create newrole "NewRole Description" --title="title" --description="description" --whitelist="1,2" --blacklist="3,4" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes

tsukid query customgov proposals
tsukid query customgov proposal 1

tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 

# check newly created role
tsukid query customgov role 3
