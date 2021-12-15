package funcs

import (
	"mmocker/utils/log"
)

// GetFunc return the func by param, and some func needn't the param.
func GetFunc(name string, param map[string]float64) Function {
	log.Logger.Infof("Load function: {%s} with param: {%v}", name, param)
	switch name {
	case "StandardLinearFunction": // y = slope * (x + offsetX) + offsetY
		return NewLinearFunctionByMap(param)
	case "DefaultLinearFunction": // y = x
		return DefaultLinearFunction()
	case "ReverseLinearFunction": // y = -x
		return ReverseLinearFunction()
	default:
		// can't find any function, return funcs.ZeroFunction
		return ZeroFunction{}
	}
}
