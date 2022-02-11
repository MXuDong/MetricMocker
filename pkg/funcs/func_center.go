package funcs

type FunctionParams struct {
	Type TypeStr `yaml:"Type"`
	Name string `yaml:"Name"`
	Params map[interface{}]interface{} `yaml:"Params"`
	KeyFunctions map[string]FunctionParams `yaml:"KeyFunctions"`
}

func Function(param FunctionParams) BaseFuncInterface {

	functionKeyFunctions := map[string]BaseFuncInterface{}
	var funcItem BaseFuncInterface

	if param.KeyFunctions != nil{
		for key, funcParams := range param.KeyFunctions {
			functionKeyFunctions[key] = Function(funcParams)
		}
	}


	switch param.Type{
	case "BaseLinearFunction":
		funcItem = &BaseLinearFunction{}
	}

	if funcItem == nil{
		return nil
	}

	funcItem.Init(param.Params)

	for key, keyFunction := range functionKeyFunctions {
		funcItem.SetKeyFunc(key, keyFunction)
	}

	return funcItem
}
