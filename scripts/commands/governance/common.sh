#!/bin/bash

tsukid tx customgov councilor claim-seat --from validator --keyring-backend=test --home=$HOME/.tsukid

tsukid tx customgov permission blacklist-permission
tsukid tx customgov permission remove-blacklisted-permission
tsukid tx customgov permission whitelist-permission
tsukid tx customgov permission remove-whitelisted-permission

# add / remove / query whitelisted permissions
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=7 --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid query customgov permissions $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)
tsukid tx customgov permission remove-whitelisted-permission --from validator --keyring-backend=test --permission=7 --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid query customgov permissions $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

tsukid tx customgov proposal assign-permission
tsukid tx customgov proposal vote

tsukid tx customgov role blacklist-permission
tsukid tx customgov role create
tsukid tx customgov role remove
tsukid tx customgov role remove-blacklist-permission
tsukid tx customgov role remove-whitelist-permission
tsukid tx customgov role whitelist-permission

# query all roles
tsukid query customgov all-roles
# query roles for an address
tsukid query customgov roles $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# querying for voters of a specific proposal
tsukid query customgov voters 1
# querying for votes of a specific proposal
tsukid query customgov votes 1
# querying for a vote of a specific propsal/voter pair
tsukid query customgov vote 1 $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# whitelist permission for claim validator
tsukid keys add lladmin --keyring-backend=test
tsukid tx bank send validator $(tsukid keys show -a lladmin --keyring-backend=test) 1000000ukex --keyring-backend=test --chain-id=testing --fees=200ukex --yes
tsukid tx customgov permission whitelist-permission --from=validator --keyring-backend=test --addr=$(tsukid keys show -a lladmin --keyring-backend=test) --permission=30 --chain-id=testing --fees=200ukex --yes
tsukid tx customgov permission whitelist-permission --from=lladmin --keyring-backend=test --addr=$(tsukid keys show -a lladmin --keyring-backend=test) --permission=2 --chain-id=testing --fees=200ukex --yes
