#!/bin/bash

rm -rf $HOME/.tsukid/

cd $HOME

tsukid init --chain-id=testing testing --home=$HOME/.tsukid
tsukid keys add validator --keyring-backend=test --home=$HOME/.tsukid
tsukid add-genesis-account $(tsukid keys show validator -a --keyring-backend=test --home=$HOME/.tsukid) 1000000000ukex,1000000000validatortoken,1000000000stake,10000000frozen  --home=$HOME/.tsukid
tsukid gentx-claim validator --keyring-backend=test --moniker="hello" --home=$HOME/.tsukid
tsukid start --home=$HOME/.tsukid

