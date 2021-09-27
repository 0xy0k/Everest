#!/bin/bash

# upsert rate by governance
## proposal
tsukid tx tokens proposal-upsert-rate --from=validator --keyring-backend=test --title="upsert rate" --description="upsert rate description" --denom="mykex" --rate="1.5" --fee_payments=true --chain-id=testing --fees=100ukex --home=$HOME/.tsukid  --yes

## query
tsukid query proposals
## vote
tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 