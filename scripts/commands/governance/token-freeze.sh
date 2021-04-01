#!/bin/bash

# create a proposal to blacklist validatortoken
tsukid tx tokens propose-update-tokens-blackwhite --is_blacklist=true --is_add=true --tokens=validatortoken --tokens=kava --from validator --chain-id=testing --keyring-backend=test --fees=100ukex --home=$HOME/.tsukid --yes
# check proposal ID
tsukid query customgov proposals
# whitelist permission to vote on proposal
tsukid tx customgov permission whitelist-permission --from validator --keyring-backend=test --permission=$PermVoteTokensWhiteBlackChangeProposal --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid --yes
# vote on proposal
tsukid tx customgov proposal vote 1 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes 
# get all votes
tsukid query customgov vote 1 $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)