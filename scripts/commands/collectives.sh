#!/bin/bash

## queries
tsukid query bank balances $(tsukid keys show -a validator --keyring-backend=test)
tsukid query customstaking validator --addr=$(tsukid keys show -a validator --keyring-backend=test)
tsukid query collectives collective userincentives
tsukid query collectives collectives
tsukid query collectives collectives-by-account
tsukid query collectives collectives-proposals

tsukid query spending pool-by-name "UserIncentivesPool"
tsukid query spending pools-by-account $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

## multistaking pool creation
tsukid tx multistaking upsert-staking-pool tsukivaloper1ak6c3jl4svl5vw5y9xu3yrq4susvkckwurn4sc --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block

## multistake and get staking tokens
tsukid tx multistaking delegate tsukivaloper1ak6c3jl4svl5vw5y9xu3yrq4susvkckwurn4sc 10000000000000ukex --from=validator --keyring-backend=test --fees=100ukex --chain-id=testing -y --broadcast-mode=block

## create dynamic spending pool
tsukid tx spending create-spending-pool --name="UserIncentivesPool" --claim-start=$(($(date -u +%s))) --claim-end=0 --claim-expiry=43200 --rates=0.1ukex --vote-quorum="33" --vote-period="60" --vote-enactment="30" --owner-roles="" --owner-accounts=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --beneficiary-roles="1" --beneficiary-role-weights="1" --beneficiary-accounts="" --beneficiary-account-weights="" --dynamic-rate=true --dynamic-rate-period=43200 --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes  --broadcast-mode=block 

## register as beneficiary
tsukid tx spending register-spending-pool-beneficiary --name="UserIncentivesPool" --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block 

## collective creation
tsukid tx collectives create-collective --collective-name="userincentives" --collective-description="description" --bonds="9000000000000v1/ukex" --deposit-any=true --deposit-roles="1" --deposit-accounts="" --owner-roles="1" --owner-accounts="" --weighted-spending-pools="UserIncentivesPool#1" --claim-start=0 --claim-period=43200 --claim-end=86400 --vote-quorum=30 --vote-period=86400 --vote-enactment=1000 --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block 

## collective contribute
tsukid tx collectives contribute-collective --collective-name="userincentives" --bonds="10000000000v1/ukex" --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block 

## collective donation
tsukid tx collectives donate-collective --collective-name="userincentives" --locking=86400 --donation="0.1" --donation-lock=false --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block 

## collective contribution withdraw
tsukid tx collectives withdraw-collective --collective-name="userincentives" --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block

## donation withdraw proposal
tsukid tx collectives proposal-send-donation --title="title" --description="description" --collective-name="userincentives" --addr="tsuki1ak6c3jl4svl5vw5y9xu3yrq4susvkckw090kg5" --amounts="100ukex" --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block

## collective update proposal
tsukid tx collectives proposal-collective-update --title="title" --description="description" --collective-name="userincentives" --collective-description="description" --collective-status="ACTIVE" --deposit-any=true --deposit-roles="1" --deposit-accounts="" --owner-roles="1" --owner-accounts="" --weighted-spending-pools="UserIncentivesPool#1" --claim-start=0 --claim-period=43200 --claim-end=86400 --vote-quorum=30 --vote-period=86400 --vote-enactment=1000 --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block 

## collective remove proposal
tsukid tx collectives proposal-remove-collective --title="title" --description="description" --collective-name="userincentives" --from=validator --chain-id=testing --fees=100ukex --keyring-backend=test --home=$HOME/.tsukid --yes --broadcast-mode=block 

tsukid query customgov proposals
tsukid query customgov proposal 1

tsukid tx customgov proposal vote 2 1 --from validator --keyring-backend=test --home=$HOME/.tsukid --chain-id=testing --fees=100ukex --yes  --broadcast-mode=block 

## Test scenarios
# - Do not allow the creation of collectives with a name that already exists
# - Ensure that the collective creator must bond a minimum of 10% `min_collective_bond` while sending create tx 
#    and that there is no minimum deposit limit for any other contributors to the collective.
# - Ensure that collective is removed and NEVER becomes active if full `min_collective_bond` is NOT reached
# - Ensure that the collective creator and all depositors can retrieve their funds if the collective is removed or
#   otherwise get them back automatically if such a dissolution event takes place during or after the creation of the collective.
# - Test withdraws in combination with locking period, verify if locking period can be extended 
#   (e.g. to 10y by sending locking tx 10 times)
# - Ensure that ANY staking derivative tokens as well as multiple staking derivative tokens 
#   can be used simultaneously to create collective but NOT any other types of tokens.
# - Ensure rewards start being claimed at `claim-start` and end at `claim-end`
# - Ensure rewards claim doesn’t occur more often than as defined by `claim-period`
# - Verify that every time rewards are distributed to the spending pool the `last-claim` property is updated
# - Verify that owners can raise and vote on proposals
# - Ensure that contributor can withdraw their tokens and trigger the collective removal proposal automatically if the token balance falls below `min_collective_bond`.
# - Ensure that collective status changes to inactive if the deposited balance falls below `min_collective_bond`. (e.g. one of the staking tokens is delisted or its value in regards to KEX changes)
# - Ensure that owners can transfer tokens from staking reward donations to any address of their choosing.
# - Verify that the spending pool module can distribute all deposited tokens to all registered beneficiaries, within `dynamic-rate-period`.
# - Create integration tests where Staking Collective and Spending Pool are interconnected and operational together.

