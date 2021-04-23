#!/bin/bash

# command
tsukid query customgov permissions $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# response
# blacklist: []
# whitelist:
# - 4
# - 3