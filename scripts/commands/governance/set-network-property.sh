#!/bin/bash

tsukid tx customgov proposal set-network-property MIN_TX_FEE 101 --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes