package functions

import (
	functionmeta "github.com/TsukiCore/tsuki/function_meta"
	tsukitypes "github.com/TsukiCore/tsuki/types"
)

// GetAllFunctions is a function to get all functions registered
func GetAllFunctions() tsukitypes.FunctionList {
	return functionmeta.GetFunctionList()
}
