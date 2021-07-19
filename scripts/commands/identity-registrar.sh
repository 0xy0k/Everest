#!/bin/bash

# create id.json file with below content
# {
#     "key1": "value1",
#     "key2": "value2"
# }

tsukid tx customgov create-identity-record --infos-file="id.json" --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid query customgov all-identity-records --log_level=debug
tsukid query customgov identity-record 1 --log_level=debug
tsukid query customgov identity-records-by-addr $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --log_level=debug
# pagination: null
# records:
# - address: tsuki1zakwshqmx92fkl7ps094u4aratk827knfl7hm2
#   date: "2021-07-06T12:31:21Z"
#   id: "1"
#   infos:
#     key1: value1
#     key2: value2
#   verifiers: []

tsukid tx customgov edit-identity-record --record-id=1 --infos-file="id.json" --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid tx customgov request-identity-record-verify --record-ids=1 --verifier=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --tip=10ukex --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes

tsukid query customgov all-identity-record-verify-requests --log_level=debug
tsukid query customgov identity-record-verify-request 1 --log_level=debug
tsukid query customgov identity-record-verify-requests-by-approver $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --log_level=debug
tsukid query customgov identity-record-verify-requests-by-requester $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --log_level=debug
# pagination: null
# verify_records:
# - address: tsuki1zakwshqmx92fkl7ps094u4aratk827knfl7hm2
#   id: "1"
#   recordIds:
#   - "1"
#   tip:
#     amount: "10"
#     denom: ukex
#   verifier: tsuki1zakwshqmx92fkl7ps094u4aratk827knfl7hm2

tsukid tx customgov approve-identity-records 1 --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes --log_level=debug

tsukid tx customgov request-identity-record-verify --record-ids=1 --verifier=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --tip=10ukex --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes
tsukid query customgov all-identity-record-verify-requests --log_level=debug
# pagination: null
# verify_records:
# - address: tsuki1zakwshqmx92fkl7ps094u4aratk827knfl7hm2
#   id: "2"
#   recordIds:
#   - "1"
#   tip:
#     amount: "10"
#     denom: ukex
#   verifier: tsuki1zakwshqmx92fkl7ps094u4aratk827knfl7hm2

tsukid tx customgov cancel-identity-records-verify-request 2 --from=validator --keyring-backend=test --home=$HOME/.tsukid --fees=100ukex --chain-id=testing --yes --log_level=debug
tsukid query customgov all-identity-record-verify-requests --log_level=debug
# pagination: null
# verify_records: []