package funcs

import "mmocker/utils/log"

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

	ModularFunctionType:     ModularFunctionInitiator,
	TimeSecondsFunctionType: TimeSecondsFunctionInitiator,
	//TimeMinutesFunctionType: TimeMinutesFunctionInitiator,
	TimeSecondsInHourFunctionType: TimeSecondsInHourFunctionInitiator,
}

func Function(param FunctionParams) BaseFuncInterface {
	log.Logger.Infof("Get function: type: [%s], params: %v", param.Type, param.Params)
	functionKeyFunctions := map[string]BaseFuncInterface{}
	var funcItem BaseFuncInterface

	if funcItemInitFunc, ok := FuncMap[param.Type]; !ok {
		return nil
	} else {
		funcItem = InitFunction(funcItemInitFunc(), param.Params)
		funcItem.SetType(param.Type)
	}

	if !funcItem.IsDerived() {
		// if function is Derived, skip set KeyFunction.
		if param.KeyFunctions != nil {
			for key, funcParams := range param.KeyFunctions {
				functionKeyFunctions[key] = Function(funcParams)
			}
		}
		for key, _ := range funcItem.KeyMap() {
			// set default KeyFunction
			// use MetadataUnitFunction as default function.
			funcItem.SetKeyFunc(key, &MetadataUnitFunction{})
		}
		for key, keyFunction := range functionKeyFunctions {
			funcItem.SetKeyFunc(key, keyFunction)
		}
	} else {
		// cover all KeyFunction
		keyMaps := funcItem.Keys()
		keyFunctions := funcItem.KeyMap()
		for key, _ := range keyMaps {
			if _, ok := keyFunctions[key]; !ok {
				// if key has no function, set default key-function.
				funcItem.SetKeyFunc(key, &MetadataUnitFunction{})
			}
		}
	}

	return funcItem
}
