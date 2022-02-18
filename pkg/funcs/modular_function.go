package funcs

import (
	"fmt"
	"math"
)

const ModularFunctionType = "ModularFunctionType"

type ModularFunction struct {
	BaseFunc
	baseExpression string `expression:"y=x%modular_unit"`
	params         map[string]interface{}
	ModularUnit    float64 `key:"modular_unit" default:"1" mean:"modular value, can't be zero."`
}

func (m ModularFunction) Type() TypeStr {
	return ModularFunctionType
}

func (m ModularFunction) Expression() string {
	return fmt.Sprintf("%s%%%f", m.KeyExpressionMap()[UnknownKey], m.ModularUnit)
}

func (m ModularFunction) Init() {
	m.BaseFunc.BaseInit()
}

func (m ModularFunction) Doc() string {
	return `
ModularFunction is a modular function, to get the value % specify value.`
}

func (m *ModularFunction) Params() map[string]interface{} {
	if m.params == nil {
		m.params = GetParamMap(m)
	}
	return m.params
}

func (m ModularFunction) KeyMap() map[string]struct{} {
	return map[string]struct{}{
		UnknownKey: {},
	}
}

func (m ModularFunction) Call(f float64) (float64, error) {
	if m.ModularUnit == 0 {
		return 0, ZeroValueError.Param("ModularFunction.modular_unit")
	}
	x, err := m.Keys()[UnknownKey].Call(f)
	if err != nil {
		return 0, err
	}

	return math.Mod(x, m.ModularUnit), nil
}
