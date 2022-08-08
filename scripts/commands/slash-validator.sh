#!/bin/bash

tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test)
tsukid query customstaking validator --addr=$(tsukid keys show -a validator --keyring-backend=test)
tsukid query multistaking pools

tsukid tx multistaking upsert-staking-pool tsukivaloper1zkka0hucf8mt8wwhqfft7pc0cgv7gk0merzmrs --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block
tsukid tx multistaking delegate tsukivaloper1zkka0hucf8mt8wwhqfft7pc0cgv7gk0merzmrs 1000000ukex --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block

# proposal to create slash validator
tsukid tx customslashing proposal-slash-validator --offender=tsukivaloper1zkka0hucf8mt8wwhqfft7pc0cgv7gk0merzmrs --staking-pool-id=1 --misbehaviour-time=1659927223 --misbehaviour-type="manual-slash" --jail-percentage=10 --colluders="" --refutation="" --title="Slash validator" --description="Slash valiator" --from=validator --chain-id=testing --keyring-backend=test  --fees=100ukex --yes --log_format=json --broadcast-mode=async --output=json --home=$HOME/.tsukid --broadcast-mode=block

# refute slash proposal
tsukid tx customslashing refute-slash-validator-proposal --refutation="refutation.com/1" --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block

# vote slash validator proposal
tsukid tx customgov proposal vote 1 1 --slash=20 --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block

# query slash validator proposals
tsukid query customslashing slash-proposals 
