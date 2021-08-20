#!/bin/bash

# create id.json file with below content
# {
#     "key1": "value1",
#     "key2": "value2"
# }

tsukid tx customgov create-identity-record --infos-file="id.json" --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
tsukid tx customgov create-identity-record --infos-json='{"key":"moniker","value":"My Moniker"}' --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid query customgov all-identity-records
tsukid query customgov identity-record 1
tsukid query customgov identity-record-by-addr $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

tsukid tx customgov edit-identity-record --record-id=1 --infos-file="id.json" --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
tsukid tx customgov edit-identity-record --record-id=1 --infos-json='{"key":"moniker","value":"My Moniker"}' --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid tx customgov request-identity-record-verify --record-ids=1 --verifier=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --tip=10ukex --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid query customgov all-identity-record-verify-requests
tsukid query customgov identity-record-verify-request 1
tsukid query customgov identity-record-verify-requests-by-approver $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)
tsukid query customgov identity-record-verify-requests-by-requester $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

tsukid tx customgov approve-identity-records 1 --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid tx customgov request-identity-record-verify --record-ids=1 --verifier=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --tip=10ukex --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
tsukid query customgov all-identity-record-verify-requests

tsukid tx customgov cancel-identity-records-verify-request 2 --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
tsukid query customgov all-identity-record-verify-requests
