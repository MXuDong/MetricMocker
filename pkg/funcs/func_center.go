package funcs

type FunctionParams struct {
	Type         TypeStr                   `yaml:"Type" json:"Type"`
	Name         string                    `yaml:"Name" json:"Name"`
	Params       map[string]interface{}    `yaml:"Params" json:"Params"`
	KeyFunctions map[string]FunctionParams `yaml:"KeyFunctions" json:"KeyFunctions"`
}
type FuncInitiator func() BaseFuncInterface

var (
	TrueP  = true
	FalseP = false
)

var FuncMap = map[TypeStr]FuncInitiator{
	// base function
	MetadataUnitFunctionType: func() BaseFuncInterface { return &MetadataUnitFunction{} },
	StartZeroFuncType:        func() BaseFuncInterface { return &StartZeroFunc{} },

	BaseLinearFunctionType:          BaseLinearFunctionInitiator,
	SingleLinearFunctionType:        SingleLinearFunctionInitiator,
	ReverseSingleLinearFunctionType: ReverseSingleLinearFunctionInitiator,

	ModularFunctionType:    ModularFunctionInitiator,
	TimeSecondFunctionType: TimeSecondFunctionInitiator,
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
		funcItem.SetType(param.Type)
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
