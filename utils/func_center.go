package utils

import "mmocker/pkg/funcs"

// GetFunc return the func by param, and some func needn't the param.
func GetFunc(name string, param map[string]float64) funcs.Function {
	switch name {
	case "StandardLinearFunction": // y = slope * (x + offsetX) + offsetY
		return funcs.NewLinearFunctionByMap(param)
	case "DefaultLinearFunction": // y = x
		return funcs.DefaultLinearFunction()
	case "ReverseLinearFunction": // y = -x
		return funcs.ReverseLinearFunction()
	default:
		// can't find any function, return funcs.ZeroFunction
		return funcs.ZeroFunction{}
	}
}
