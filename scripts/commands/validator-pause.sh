#!/bin/bash

tsukid tx customslashing pause --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block 
