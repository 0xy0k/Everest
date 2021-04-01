#!/bin/bash

# upsert alias by governance
## proposal
tsukid tx tokens proposal-upsert-alias --from=validator --keyring-backend=test --symbol="ETH" --name="Ethereum" --icon="myiconurl" --decimals=6 --denoms="finney" --chain-id=testing --fees=100ukex --home=$HOME/.tsukid  --yes
## query
tsukid query proposals
## vote
tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 