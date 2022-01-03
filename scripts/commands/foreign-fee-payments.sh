#!/bin/bash

# register stake token as 1ukex=100stake
PermUpsertTokenRate=8
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermUpsertTokenRate --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="stake" --rate="0.01" --fee_payments=true --chain-id=testing --fees=100ukex --home=$HOME/.tsukid  --yes
tsukid query tokens rate stake
# try to spend stake token as fee
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="valstake" --rate="0.01" --fee_payments=true --chain-id=testing --fees=10000stake --home=$HOME/.tsukid  --yes
# smaller amount of fee in foreign currency
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="valstake" --rate="0.02" --fee_payments=true --chain-id=testing --fees=1000stake --home=$HOME/.tsukid  --yes
# try to spend unregistered token (validatortoken) as fee
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="valstake" --rate="0.03" --fee_payments=true --chain-id=testing --fees=1000validatortoken --home=$HOME/.tsukid  --yes