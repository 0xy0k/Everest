#!/bin/bash

rm -rf $HOME/.tsukid/
rm -rf $HOME/.tsukicli/

cd $HOME

tsukid init --chain-id=testing testing
tsukicli keys add validator
tsukid add-genesis-account $(tsukicli keys show validator -a) 1000000000stake,1000000000validatortoken
tsukid gentx --name validator
tsukid collect-gentxs
tsukid start
