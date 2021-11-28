#!/bin/bash

tsukid tx customgov proposal upsert-data-registry "some_record" 0x0000 "https://github.githubassets.com/images/modules/notifications/inbox-zero.svg" "image/svg+xml" 31088 --keyring-backend=test --from=validator --home=$HOME/.tsukid --title="upsert-data-registry" --description="upsert data" --chain-id=testing --yes --fees=100ukex
tsukid tx customgov proposal vote 1 1 --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --yes --fees=100ukex
tsukid query customgov proposals
tsukid query customgov data-registry-keys
tsukid query customgov data-registry "some_record"