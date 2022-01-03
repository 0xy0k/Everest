#!/bin/bash

# e.g. tsukid query customgov execution-fee <msg_type>

# command
tsukid query customgov execution-fee "B"
# response
# fee:
#   default_parameters: "0"
#   execution_fee: "10"
#   failure_fee: "1"
#   name: ABC
#   timeout: "10"
#   transaction_type: B

# genesis fee configuration test
tsukid query customgov execution-fee "A"
# response
# fee:
#   default_parameters: "0"
#   execution_fee: "10"
#   failure_fee: "1"
#   name: Claim Validator Seat
#   timeout: "10"
#   transaction_type: A