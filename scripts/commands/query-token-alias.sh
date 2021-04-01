#!/bin/bash

# command
tsukid query tokens alias KEX
# response
# allowed_vote_types:
# - "yes"
# - "no"
# decimals: 6
# denoms:
# - ukex
# enactment: 0
# expiration: 0
# icon: myiconurl
# name: Tsuki
# status: undefined
# symbol: KEX


# command
tsukid query tokens alias KE
# response
# Error: KE symbol does not exist

# command
tsukid query tokens all-aliases --chain-id=testing --home=$HOME/.tsukid
# response
# data:
# - allowed_vote_types:
#   - "yes"
#   - "no"
#   decimals: 6
#   denoms:
#   - ukex
#   enactment: 0
#   expiration: 0
#   icon: myiconurl
#   name: Tsuki
#   status: undefined
#   symbol: KEX

# command
tsukid query tokens aliases-by-denom ukex --chain-id=testing --home=$HOME/.tsukid
# response
# data:
#   ukex:
#     allowed_vote_types:
#     - "yes"
#     - "no"
#     decimals: 6
#     denoms:
#     - ukex
#     enactment: 0
#     expiration: 0
#     icon: myiconurl
#     name: Tsuki
#     status: undefined
#     symbol: KEX