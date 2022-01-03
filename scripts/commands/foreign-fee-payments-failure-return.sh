#!/bin/bash

# register stake token as 1ukex=100stake
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermUpsertTokenRate --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
tsukid tx tokens upsert-rate --from validator --keyring-backend=test --denom="stake" --rate="0.01" --fee_payments=true --chain-id=testing --fees=100ukex --home=$HOME/.tsukid  --yes
tsukid query tokens rate stake

# set execution fee and failure fee for upsert-rate transaction
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermChangeTxFee --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes

# set execution_fee=1000 failure_fee=5000
tsukid tx customgov set-execution-fee --from validator --execution_name="upsert-token-alias" --transaction_type="upsert-token-alias" --execution_fee=1000 --failure_fee=5000 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes

# set execution_fee=5000 failure_fee=1000
tsukid tx customgov set-execution-fee --from validator --execution_name="upsert-token-alias" --transaction_type="upsert-token-alias" --execution_fee=5000 --failure_fee=1000 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes

# check current balance
tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# try upsert-token-alias failure in foreign currency
tsukid tx tokens upsert-alias --from validator --keyring-backend=test --expiration=0 --enactment=0 --allowed_vote_types=0,1 --symbol="ETH" --name="Ethereum" --icon="myiconurl" --decimals=6 --denoms="finney" --chain-id=testing --fees=500000stake --home=$HOME/.tsukid  --yes
# set permission for this execution
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermUpsertTokenAlias --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=10000stake --home=$HOME/.tsukid --yes
# try upsert-token-alias success in foreign currency
tsukid tx tokens upsert-alias --from validator --keyring-backend=test --expiration=0 --enactment=0 --allowed_vote_types=0,1 --symbol="ETH" --name="Ethereum" --icon="myiconurl" --decimals=6 --denoms="finney" --chain-id=testing --fees=500000stake --home=$HOME/.tsukid  --yes