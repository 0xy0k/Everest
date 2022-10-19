#!/bin/bash

# councilor management commands
tsukid tx customgov councilor claim-seat --moniker="seedvalidator" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex
tsukid tx customgov councilor pause --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex
tsukid tx customgov councilor unpause --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex
tsukid tx customgov councilor activate --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex

# councilor proposal commands
tsukid tx customgov proposal proposal-reset-whole-councilor-rank --title="Title" --description="Description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex

# councilor queries
tsukid query customgov councilors
tsukid query customgov non-councilors
tsukid query customgov whitelisted-permission-addresses 1
tsukid query customgov blacklisted-permission-addresses 1
tsukid query customgov whitelisted-role-addresses 1
tsukid query customgov blacklisted-role-addresses 1
