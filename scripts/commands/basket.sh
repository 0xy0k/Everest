#!/bin/bash

# queries
tsukid query basket token-basket-by-id 1
tsukid query basket token-basket-by-denom b1/usd
tsukid query basket token-baskets "" false
tsukid query basket historical-burns 1
tsukid query basket historical-mints 1
tsukid query basket historical-swaps 1

# transactions
tsukid tx basket mint-basket-tokens 1 1000stake --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx basket mint-basket-tokens 1 10000stake,100000ukex,1000000ueth --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block

tsukid tx basket burn-basket-tokens 1 10000b1/usd --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx basket swap-basket-tokens 1 100000stake ueth --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx basket basket-claim-rewards 1000b1/usd --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx basket disable-basket-deposits 1 --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx basket disable-basket-swaps 1 --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx basket disable-basket-withdraws 1 --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block

# proposals
tsukid keys add receiver --keyring-backend=test
tsukid tx basket proposal-create-basket --title="title" --description="description" --basket-suffix="usd" --basket-description="usd stable coins collection" --swap-fee="0.001" --slippage-fee-min="0.001" --tokens-cap="1.0" --limits-period="86400" --mints-min="10000" --mints-max="1000000" --mints-disabled=false --burns-min="10000" --burns-max="1000000" --burns-disabled=false --swaps-min="10000" --swaps-max="10000" --swaps-disabled=false --basket-tokens="stake#10#false#false#false,ukex#1#false#false#false,ueth#0.1#false#false#true" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx basket proposal-edit-basket --title="title" --description="description" --basket-id="1" --basket-suffix="usd" --basket-description="usd stable coins collection" --swap-fee="0.001" --slippage-fee-min="0.001" --tokens-cap="1.0" --limits-period="86400" --mints-min="10000" --mints-max="1000000" --mints-disabled=false --burns-min="10000" --burns-max="1000000" --burns-disabled=false --swaps-min="1000" --swaps-max="1000000" --swaps-disabled=false --basket-tokens="stake#10#true#true#true,ukex#1#true#true#true,ueth#0.1#true#true#true" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx basket proposal-edit-basket --title="title" --description="description" --basket-id="1" --basket-suffix="usd" --basket-description="usd stable coins collection" --swap-fee="0.001" --slippage-fee-min="0.001" --tokens-cap="0.9" --limits-period="86400" --mints-min="10000" --mints-max="1000000" --mints-disabled=false --burns-min="10000" --burns-max="1000000" --burns-disabled=false --swaps-min="1000" --swaps-max="1000000" --swaps-disabled=false --basket-tokens="stake#10#false#false#false,ukex#1#false#false#false,ueth#0.1#false#false#false" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx basket proposal-basket-withdraw-surplus 1 $(tsukid keys show -a receiver --keyring-backend=test) --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx basket proposal-basket-withdraw-surplus 1,2 $(tsukid keys show -a receiver --keyring-backend=test) --title="title" --description="description" --from=validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block

tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx customgov proposal vote 2 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block
tsukid tx customgov proposal vote 3 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes --broadcast-mode=block

tsukid query customgov proposals
tsukid query bank balances $(tsukid keys show -a receiver --keyring-backend=test)
tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test)