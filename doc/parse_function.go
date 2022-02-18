package doc

import "mmocker/pkg/funcs"

type FunctionDescribe struct {
	FunctionName string
	Keys         map[string]funcs.FieldDescribe
	Doc          string
	Expression   string
}

func GetFunctionDescribe(funcInterface funcs.BaseFuncInterface) FunctionDescribe {
	return FunctionDescribe{
		FunctionName: string(funcInterface.Type()),
		Doc:          funcInterface.Doc(),
		Expression:   funcs.GetExpression(funcInterface),
		Keys:         funcs.GetParamFields(funcInterface),
	}
}
