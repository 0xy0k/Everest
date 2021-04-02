#!/bin/bash

# query validator account by address
tsukid query validator --addr $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)