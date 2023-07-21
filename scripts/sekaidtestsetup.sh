#!/bin/bash

rm -rf $HOME/.tsukid/

cd $HOME

tsukid init --default-denom="ukex" --bech32-prefix="tsuki" --chain-id=testing testing --home=$HOME/.tsukid
tsukid keys add validator --keyring-backend=test --home=$HOME/.tsukid
tsukid add-genesis-account $(tsukid keys show validator -a --keyring-backend=test --home=$HOME/.tsukid) 1000000000000000ukex,1000000000ubtc,1000000000ueth,1000000000validatortoken,1000000000stake,10000000frozen,10000000samolean  --home=$HOME/.tsukid
tsukid gentx-claim validator --keyring-backend=test --moniker="hello" --home=$HOME/.tsukid

cat $HOME/.tsukid/config/genesis.json | jq '.app_state["customgov"]["network_properties"]["minimum_proposal_end_time"]="30"' > $HOME/.tsukid/config/tmp_genesis.json && mv $HOME/.tsukid/config/tmp_genesis.json $HOME/.tsukid/config/genesis.json
cat $HOME/.tsukid/config/genesis.json | jq '.app_state["customgov"]["network_properties"]["proposal_enactment_time"]="10"' > $HOME/.tsukid/config/tmp_genesis.json && mv $HOME/.tsukid/config/tmp_genesis.json $HOME/.tsukid/config/genesis.json

tsukid start --home=$HOME/.tsukid
