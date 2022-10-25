#!/bin/bash

# councilor management commands
tsukid tx customgov councilor claim-seat --moniker="seedvalidator" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex
tsukid tx customgov councilor pause --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex
tsukid tx customgov councilor unpause --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex
tsukid tx customgov councilor activate --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex

# councilor proposal commands
tsukid tx customgov proposal proposal-reset-whole-councilor-rank --title="Title" --description="Description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex
tsukid tx customgov proposal proposal-jail-councilor $(tsukid keys show -a councilor1 --keyring-backend=test --home=$HOME/.tsukid) --title="Title" --description="Description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex
tsukid tx customgov proposal vote 1 1 --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex

# councilor queries
tsukid query customgov councilors
tsukid query customgov non-councilors
tsukid query customgov whitelisted-permission-addresses 1
tsukid query customgov blacklisted-permission-addresses 1
tsukid query customgov whitelisted-role-addresses 1
tsukid query customgov blacklisted-role-addresses 1

# check waiting councilor creation on permission assign
tsukid keys add councilor1 --keyring-backend=test
PermClaimCouncilor=3
tsukid tx customgov permission whitelist --from validator --keyring-backend=test --permission=$PermClaimCouncilor --addr=$(tsukid keys show -a councilor1 --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes --broadcast-mode=block

# check waiting councilor creation on role assign
tsukid keys add councilor2 --keyring-backend=test
RoleSudo=1
tsukid tx customgov role assign $RoleSudo --addr=$(tsukid keys show -a councilor2 --keyring-backend=test --home=$HOME/.tsukid) --from=validator --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes --broadcast-mode=block

# try claim councilor
tsukid tx bank send validator $(tsukid keys show -a councilor2 --keyring-backend=test --home=$HOME/.tsukid) 100000ukex --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes --broadcast-mode=block
tsukid tx customgov councilor claim-seat --moniker="councilor2" --from=councilor2 --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex

# open a proposal
tsukid tx customgov proposal proposal-reset-whole-councilor-rank --title="Title" --description="Description" --from=councilor2 --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex

# try voting twice and see rank changes
tsukid tx customgov proposal vote 1 1 --from=councilor2 --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex
tsukid tx customgov proposal vote 1 1 --from=councilor2 --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing -y --broadcast-mode=block --fees=100ukex

# moniker update on councilor
tsukid tx customgov register-identity-records --infos-json='{"moniker":"My Moniker"}' --from=councilor1 --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes --broadcast-mode=block
tsukid tx customgov register-identity-records --infos-json='{"moniker":"My Moniker2"}' --from=councilor1 --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes --broadcast-mode=block