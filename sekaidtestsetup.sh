#!/bin/bash

rm -rf $HOME/.tsukid/

cd $HOME

tsukid init --chain-id=testing testing
tsukid keys add validator --keyring-backend=test
tsukid add-genesis-account $(tsukid keys show validator -a --keyring-backend=test) 1000000000ukex,1000000000validatortoken,1000000000stake --keyring-backend=test
tsukid gentx-claim validator --keyring-backend=test --moniker="hello"
tsukid start
