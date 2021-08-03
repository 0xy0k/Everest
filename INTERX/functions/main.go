package functions

import (
	functionmeta "github.com/TsukiCore/tsuki/function_meta"
	tsukitypes "github.com/TsukiCore/tsuki/types"
)

// GetTsukiFunctions is a function to get all tsuki functions registered
func GetTsukiFunctions() tsukitypes.FunctionList {
	return functionmeta.GetFunctionList()
}

// GetInterxMetadata is a function to get all interx functions registered
func GetInterxMetadata() InterxMetadata {
	return interxMetadata
}
