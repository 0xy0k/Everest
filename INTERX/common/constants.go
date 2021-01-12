package common

const (
	QueryAccounts          = "/api/cosmos/auth/accounts/{address}"
	QueryTotalSupply       = "/api/cosmos/bank/supply"
	QueryBalances          = "/api/cosmos/bank/balances"
	PostTransaction        = "/api/cosmos/txs"
	QueryTransactionHash   = "/api/cosmos/txs"
	EncodeTransaction      = "/api/cosmos/txs/encode"
	FaucetRequestURL       = "/api/faucet"
	QueryRPCMethods        = "/api/rpc_methods"
	QueryInterxFunctions   = "/api/metadata"
	QueryStatus            = "/api/status"
	QueryWithdraws         = "/api/withdraws"
	QueryDeposits          = "/api/deposits"
	QueryDataReferenceKeys = "/api/tsuki/gov/data_keys"
	QueryDataReference     = "/api/tsuki/gov/data"
	QueryTsukiFunctions     = "/api/tsuki/metadata"
	Download               = "/download"
	DataReferenceRegistry  = "DRR"
	QueryValidators        = "/api/validators"
)
