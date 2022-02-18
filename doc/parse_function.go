package doc

import "mmocker/pkg/funcs"

type FunctionDescribe struct {
	FunctionType string
	FunctionName string
	Keys         map[string]funcs.FieldDescribe
	Doc          string
	Expression   string
}

func GetFunctionDescribe(funcInterface funcs.BaseFuncInterface) FunctionDescribe {
	return FunctionDescribe{
		FunctionType: string(funcInterface.Type()),
		FunctionName: funcs.GetFunctionName(funcInterface),
		Doc:          funcInterface.Doc(),
		Expression:   funcs.GetExpression(funcInterface),
		Keys:         funcs.GetParamFields(funcInterface),
	}
}
