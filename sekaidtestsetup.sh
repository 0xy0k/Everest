#!/bin/bash

rm -rf $HOME/.tsukid/

cd $HOME

tsukid init --chain-id=testing testing --home=/Users/jgimeno/.tsukid
tsukid keys add validator
tsukid add-genesis-account $(tsukid keys show validator -a) 1000000000stake,1000000000validatortoken
#tsukid gentx validator --chain-id=testing --offline
#tsukid collect-gentxs

tsukid gentx-claim validator
tsukid start --home=/Users/jgimeno/.tsukid
