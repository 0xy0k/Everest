#!/bin/bash

# block_time modification
# sed -i '' 's/timeout_commit = "5s"/timeout_commit = "20s"/g' $HOME/.tsukid/config/config.toml

tsukid tx customgov proposal set-poor-network-msgs AAA --description="" --from=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=2000000ukex --generate-only > AAA_proposal.json
tsukid tx customgov proposal set-poor-network-msgs BBB --description="" --from=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=3000000ukex --generate-only > BBB_proposal.json
tsukid tx customgov proposal set-poor-network-msgs CCC --description="" --from=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=4000000ukex --generate-only > CCC_proposal.json
tsukid tx customgov proposal set-poor-network-msgs DDD --description="" --from=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=1000000ukex --generate-only > DDD_proposal.json

tsukid tx sign AAA_proposal.json --log_level=debug --chain-id=testing --keyring-backend=test --from=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --home=$HOME/.tsukid > signed_AAA.json
tsukid tx sign BBB_proposal.json --log_level=debug --chain-id=testing --keyring-backend=test --from=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --home=$HOME/.tsukid > signed_BBB.json
tsukid tx sign CCC_proposal.json --log_level=debug --chain-id=testing --keyring-backend=test --from=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --home=$HOME/.tsukid > signed_CCC.json
tsukid tx sign DDD_proposal.json --log_level=debug --chain-id=testing --keyring-backend=test --from=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --home=$HOME/.tsukid > signed_DDD.json

tsukid tx broadcast signed_DDD.json --log_level=debug --broadcast-mode=async
tsukid tx broadcast signed_BBB.json --log_level=debug --broadcast-mode=async
tsukid tx broadcast signed_CCC.json --log_level=debug --broadcast-mode=async
tsukid tx broadcast signed_AAA.json --log_level=debug --broadcast-mode=async

sleep 20

tsukid query customgov proposals --log_level=debug