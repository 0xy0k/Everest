package functions

import (
	middleware "github.com/TsukiCore/tsuki/middleware"
	tsukitypes "github.com/TsukiCore/tsuki/types"
)

// GetAllFunctions is a function to get all functions registered
func GetAllFunctions() tsukitypes.FunctionList {
	return middleware.GetFunctionList()
}
