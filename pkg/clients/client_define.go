package clients

import "mmocker/pkg/common"

type BaseClientInterface interface {
	Init(param map[interface{}]interface{})
	InitP(param map[interface{}]interface{}) BaseClientInterface
	Push(processorName string, result map[string]common.FunctionResult)
}
