package funcs

type FunctionParams struct {
	Type         TypeStr                   `yaml:"Type" json:"Type"`
	Name         string                    `yaml:"Name" json:"Name"`
	Params       map[string]interface{}    `yaml:"Params" json:"Params"`
	KeyFunctions map[string]FunctionParams `yaml:"KeyFunctions" json:"KeyFunctions"`
}

var FuncMap map[TypeStr]func() BaseFuncInterface = map[TypeStr]func() BaseFuncInterface{
	"MetadataUnitFunction": func() BaseFuncInterface { return &MetadataUnitFunction{} },
	BaseLinearFunctionType: func() BaseFuncInterface { return &BaseLinearFunction{} },
	StartZeroFuncType:      func() BaseFuncInterface { return &StartZeroFunc{} },
	ModularFunctionType:    func() BaseFuncInterface { return &ModularFunction{} },
}

func Function(param FunctionParams) BaseFuncInterface {

	functionKeyFunctions := map[string]BaseFuncInterface{}
	var funcItem BaseFuncInterface

	if param.KeyFunctions != nil {
		for key, funcParams := range param.KeyFunctions {
			functionKeyFunctions[key] = Function(funcParams)
		}
	}
	if funcItemInitFunc, ok := FuncMap[param.Type]; !ok {
		return nil
	} else {
		funcItem = InitFunction(funcItemInitFunc(), param.Params)
	}

	if funcItem == nil {
		return nil
	}

	//funcItem.Init(param.Params)

	for key, _ := range funcItem.KeyMap() {
		// use MetadataUnitFunction as default function.
		funcItem.SetKeyFunc(key, &MetadataUnitFunction{})
	}

	for key, keyFunction := range functionKeyFunctions {
		funcItem.SetKeyFunc(key, keyFunction)
	}
	return funcItem
}
