package clients

import "mmocker/pkg/common"

type BaseClientInterface interface {
	Init(param map[string]interface{})
	InitP(param map[string]interface{}) BaseClientInterface
	Push(processorName string, result map[string]common.FunctionResult)
}
