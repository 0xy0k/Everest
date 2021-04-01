#!/bin/bash

tsukid tx customgov councilor claim-seat --from validator --keyring-backend=test --home=$HOME/.tsukid

tsukid tx customgov permission blacklist-permission
tsukid tx customgov permission whitelist-permission

tsukid tx customgov proposal assign-permission
tsukid tx customgov proposal vote

tsukid tx customgov role blacklist-permission
tsukid tx customgov role create
tsukid tx customgov role remove
tsukid tx customgov role remove-blacklist-permission
tsukid tx customgov role remove-whitelist-permission
tsukid tx customgov role whitelist-permission

# querying for voters of a specific proposal
tsukid query customgov voters 1
# querying for votes of a specific proposal
tsukid query customgov votes 1
# querying for a vote of a specific propsal/voter pair
tsukid query customgov vote 1 $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)