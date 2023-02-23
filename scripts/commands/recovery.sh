#!/bin/bash

# command to generate recovery secret
tsukid tx recovery generate-recovery-secret 10a0fbe01030000122300000000000
NONCE=00
PROOF=29eeb09666c4792c314e631063932621f573cbb07af7274657d1314e1892eb93
CHALLENGE=220c7e47a53ef4c2161035308d4fdc52f619e54f8a657b208ba3708139fdc03d

# commmand to create an account
tsukid keys add recovery --keyring-backend=test
RECOVERY=$(tsukid keys show -a recovery --keyring-backend=test)

# commands to query
tsukid query recovery recovery-record $(tsukid keys show -a validator --keyring-backend=test)

# command to register recovery secret
tsukid tx recovery register-recovery-secret $CHALLENGE $NONCE $PROOF --from=validator --keyring-backend=test --chain-id=testing --fees=1000ukex -y --home=$HOME/.tsukid --broadcast-mode=block

# command to rotate
tsukid tx recovery rotate-recovery-address $RECOVERY $PROOF --from=validator --chain-id=testing --keyring-backend=test --fees=1000ukex -y --home=$HOME/.tsukid --yes --broadcast-mode=block

# command to check validators after rotation
tsukid query customstaking validators --moniker="" --addr="" --val-addr="" --moniker="" --status="" --pubkey="" --proposer=""

# upsert staking pool before rotation
tsukid tx multistaking upsert-staking-pool tsukivaloper18ka9xpvwh75sgldgke69jmxsnkhjm0wa3ns9xa --commission=0.5 --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block

# recovery token queries
tsukid query recovery rr-holder-rewards $(tsukid keys show -a validator --keyring-backend=test)
tsukid query recovery rr-holders rr/hello
tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test)

# recovery token txs
tsukid keys add recovery --keyring-backend=test
tsukid keys add recovery2 --keyring-backend=test
RECOVERY=$(tsukid keys show -a recovery --keyring-backend=test)
RECOVERY2=$(tsukid keys show -a recovery2 --keyring-backend=test)
VALIDATOR=$(tsukid keys show -a validator --keyring-backend=test)

tsukid tx recovery issue-recovery-tokens --from=validator --keyring-backend=test --chain-id=testing --fees=1000ukex -y --home=$HOME/.tsukid --broadcast-mode=block
tsukid tx recovery burn-recovery-tokens 10000000rr/hello --from=validator --keyring-backend=test --chain-id=testing --fees=1000ukex -y --home=$HOME/.tsukid --broadcast-mode=block
tsukid tx recovery register-rrtoken-holder --from=validator --keyring-backend=test --chain-id=testing --fees=1000ukex -y --home=$HOME/.tsukid --broadcast-mode=block
tsukid tx recovery claim-rrtoken-rewards --from=validator --keyring-backend=test --chain-id=testing --fees=1000ukex -y --home=$HOME/.tsukid --broadcast-mode=block
tsukid tx recovery rotate-validator-by-half-rr-holder $VALIDATOR $RECOVERY --from=validator --keyring-backend=test --chain-id=testing --fees=1000ukex -y --home=$HOME/.tsukid --broadcast-mode=block

# upsert-staking-pool after recovery
tsukid tx bank send validator $RECOVERY 10000000000000rr/hello,100000ukex --from=validator --keyring-backend=test --chain-id=testing --fees=1000ukex -y --home=$HOME/.tsukid --broadcast-mode=block

tsukid tx multistaking upsert-staking-pool tsukivaloper1zwyxw66aw0fd3xv9e9ewyk58c0qdcn5synp3dt --commission=0.5 --from=recovery --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx recovery rotate-validator-by-half-rr-holder $RECOVERY $VALIDATOR --from=recovery --keyring-backend=test --chain-id=testing --fees=1000ukex -y --home=$HOME/.tsukid --broadcast-mode=block
tsukid tx recovery rotate-validator-by-half-rr-holder $RECOVERY $RECOVERY2 --from=recovery --keyring-backend=test --chain-id=testing --fees=1000ukex -y --home=$HOME/.tsukid --broadcast-mode=block
