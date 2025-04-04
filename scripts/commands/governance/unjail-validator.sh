#!/bin/bash

# make proposal to unjail validator from jailed_validator
tsukid tx customstaking proposal unjail-validator hash reference --from=jailed_validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes

# vote on unjail validator proposal
tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes

# proposal for jail max time - max to 1440min = 1d
tsukid tx customgov proposal set-network-property UNJAIL_MAX_TIME 1440  --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes