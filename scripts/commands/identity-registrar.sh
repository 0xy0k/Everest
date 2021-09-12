#!/bin/bash

# create id.json file with below content
# {
#     "key1": "value1",
#     "key2": "value2"
# }

tsukid tx customgov register-identity-records --infos-file="id.json" --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
tsukid tx customgov register-identity-records --infos-json='{"moniker":"My Moniker","social":"My Social"}' --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid query customgov all-identity-records
tsukid query customgov identity-record 1
tsukid query customgov identity-records-by-addr $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# query by specific keys
tsukid query customgov identity-records-by-addr $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --keys="moniker"

tsukid tx customgov delete-identity-records --keys="moniker" --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid keys add test --keyring-backend=test --home=$HOME/.tsukid
tsukid tx customgov request-identity-record-verify --record-ids=1 --verifier=$(tsukid keys show -a test --keyring-backend=test --home=$HOME/.tsukid) --tip=200ukex --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
tsukid tx customgov request-identity-record-verify --record-ids=2 --verifier=$(tsukid keys show -a test --keyring-backend=test --home=$HOME/.tsukid) --tip=200ukex --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid query customgov all-identity-record-verify-requests
tsukid query customgov identity-record-verify-request 1

tsukid query customgov identity-record-verify-requests-by-requester $(tsukid keys show -a test --keyring-backend=test --home=$HOME/.tsukid)
tsukid query customgov identity-record-verify-requests-by-requester $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

tsukid query customgov identity-record-verify-requests-by-approver $(tsukid keys show -a test --keyring-backend=test --home=$HOME/.tsukid)
tsukid query customgov identity-record-verify-requests-by-approver $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

tsukid tx customgov handle-identity-records-verify-request 1 --from=validator --approve=true --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
tsukid tx customgov handle-identity-records-verify-request 2 --from=validator --approve=false --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid tx customgov request-identity-record-verify --record-ids=1 --verifier=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --tip=200ukex --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
tsukid query customgov all-identity-record-verify-requests

tsukid tx customgov cancel-identity-records-verify-request 2 --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
tsukid query customgov all-identity-record-verify-requests
