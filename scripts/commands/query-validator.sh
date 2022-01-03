#!/bin/bash

# query validator account by address
tsukid query customstaking validator --addr $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# query all validators with specific filter
tsukid query customstaking validators --moniker="" --addr="" --val-addr="" --moniker="" --status="" --pubkey="" --proposer=""
