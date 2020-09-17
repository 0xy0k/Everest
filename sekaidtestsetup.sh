#!/bin/bash

rm -rf $HOME/.tsukid/

cd $HOME

tsukid init --chain-id=testing testing --home=/Users/jgimeno/.tsukid
tsukid keys add validator --keyring-backend=test --home=/Users/jgimeno/.tsukid
tsukid add-genesis-account $(tsukid keys show validator -a --home=/Users/jgimeno/.tsukid --keyring-backend=test) 1000000000stake,1000000000validatortoken  --home=/Users/jgimeno/.tsukid
tsukid gentx-claim validator --keyring-backend=test --moniker="hello" --home=/Users/jgimeno/.tsukid
tsukid start --home=/Users/jgimeno/.tsukid
