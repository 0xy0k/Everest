#!/bin/bash

rm -rf $HOME/.tsukid/
rm -rf $HOME/.tsukicli/

cd $HOME

tsukid init --chain-id=testing testing
tsukid keys add validator
tsukid add-genesis-account $(tsukid keys show validator -a) 1000000000stake,1000000000validatortoken
tsukid gentx validator --chain-id=testing --offline
tsukid collect-gentxs
tsukid start
