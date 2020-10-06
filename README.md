# tsuki
Tsuki Hub

## Set ChangeTxFee permission
```sh
# command to set changeTxFee permission
tsukid tx customgov set-whitelist-permissions --from validator --keyring-backend=test --permission=4 --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid
# good response
"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"whitelist-permissions\"}]}]}]"
```

## Query permission of an address
```sh
# command
tsukid query customgov permissions $(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid)

# response
blacklist: []
whitelist:
- 4
- 3
```
## Set network properties
```sh

# command with fee set
tsukid tx customgov set-network-properties --from validator --min_tx_fee="2" --max_tx_fee="20000" --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid

# no error response
"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"set-network-properties\"}]}]}]"

# response when not enough permissions to change tx fee
"failed to execute message; message index: 0: PermChangeTxFee: not enough permissions"

# command without fee set
tsukid tx customgov set-network-properties --from validator --min_tx_fee="2" --max_tx_fee="20000" --keyring-backend=test --chain-id=testing --home=$HOME/.tsukid

# response
"fee out of range [1, 10000]: invalid request"

```
## Query network properties
```sh
# command
tsukid query customgov network-properties

# response
properties:
  max_tx_fee: "10000"
  min_tx_fee: "1"
```

## Set Execution Fee
```sh
# command
tsukid tx customgov set-execution-fee --from validator --execution_name="ABC" --transaction_type="B" --execution_fee=10 --failure_fee=1 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=10ukex --home=$HOME/.tsukid

# response
"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"set-execution-fee\"}]}]}]"
```

## Set execution fee validation test
```sh
# command for setting execution fee
tsukid tx customgov set-execution-fee --from validator --execution_name="set-network-properties" --transaction_type="B" --execution_fee=10000 --failure_fee=1000 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid

Here, the value should be looked at is `--execution_name="set-network-properties"`, `--execution_fee=10000` and `--failure_fee=1000`.

# check execution fee validation
tsukid tx customgov set-network-properties --from validator --min_tx_fee="2" --max_tx_fee="20000" --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid
# response
"fee is less than failure fee 1000: invalid request"

Here, the value should be looked at is `"fee is less than failure fee 1000: invalid request"`.
In this case, issue is found on ante step and fee is not being paid at all.

# preparation for networks (v1) failure=1000, execution=10000
tsukid tx customgov set-whitelist-permissions --from validator --keyring-backend=test --permission=4 --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid <<< y
tsukid tx customgov set-execution-fee --from validator --execution_name="set-network-properties" --transaction_type="B" --execution_fee=10000 --failure_fee=1000 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid <<< y

# preparation for networks (v2) failure=1000, execution=500
tsukid tx customgov set-whitelist-permissions --from validator --keyring-backend=test --permission=4 --addr=$(tsukid keys show -a validator --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid <<< y
tsukid tx customgov set-execution-fee --from validator --execution_name="set-network-properties" --transaction_type="B" --execution_fee=500 --failure_fee=1000 --timeout=10 default_parameters=0 --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid <<< y

# init user1 with 100000ukex
tsukid keys add user1 --keyring-backend=test --home=$HOME/.tsukid
tsukid tx bank send validator $(tsukid keys show -a user1 --keyring-backend=test --home=$HOME/.tsukid) 100000ukex --keyring-backend=test --chain-id=testing --fees=100ukex --home=$HOME/.tsukid <<< y
tsukid query bank balances $(tsukid keys show -a user1 --keyring-backend=test --home=$HOME/.tsukid) <<< y

# try changing set-network-properties with user1 that does not have ChangeTxFee permission
tsukid tx customgov set-network-properties --from user1 --min_tx_fee="2" --max_tx_fee="25000" --keyring-backend=test --chain-id=testing --fees=1000ukex --home=$HOME/.tsukid <<< y
# this should fail and balance should be (previousBalance - failureFee)
tsukid query bank balances $(tsukid keys show -a user1 --keyring-backend=test --home=$HOME/.tsukid)

# whitelist user1's permission for ChangeTxFee and try again
tsukid tx customgov set-whitelist-permissions --from validator --keyring-backend=test --permission=4 --addr=$(tsukid keys show -a user1 --keyring-backend=test --home=$HOME/.tsukid) --chain-id=testing --fees=100ukex --home=$HOME/.tsukid <<< y
tsukid tx customgov set-network-properties --from user1 --min_tx_fee="2" --max_tx_fee="25000" --keyring-backend=test --chain-id=testing --fees=1000ukex --home=$HOME/.tsukid <<< y
# this should fail and balance should be (previousBalance - successFee)
tsukid query bank balances $(tsukid keys show -a user1 --keyring-backend=test --home=$HOME/.tsukid)
```

## Query execution fee
```sh
tsukid query customgov execution-fee <msg_type>
# command
tsukid query customgov execution-fee ABC
# response
fee:
  default_parameters: "0"
  execution_fee: "10"
  failure_fee: "1"
  name: ABC
  timeout: "10"
  transaction_type: B

# genesis fee configuration test
tsukid query customgov execution-fee "Claim Validator Seat"
fee:
  default_parameters: "0"
  execution_fee: "10"
  failure_fee: "1"
  name: Claim Validator Seat
  timeout: "10"
  transaction_type: A
```

---
`dev` branch
