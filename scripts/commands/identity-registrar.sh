#!/bin/bash

# create id.json file with below content
# {
#     "key1": "value1",
#     "key2": "value2"
# }

tsukid tx customgov create-identity-record --infos-file="id.json" --timestamp=1625574681 --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
