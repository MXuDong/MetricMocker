package funcs

import (
	"fmt"
	"math"
)

type ModularFunction struct {
	BaseFunc
	baseExpression string `expression:"y=(x+offsetX)%modular_unit+offsetY"`
	params         map[string]interface{}
	ModularUnit    float64 `key:"modular_unit" default:"1" mean:"modular value, can't be zero."`
	OffsetX        float64 `key:"offsetX" default:"0" mean:"offset of x"`
	OffsetY        float64 `key:"offsetY" default:"0" mean:"offset of y"`
}

func (m ModularFunction) Expression() string {
	return fmt.Sprintf("(%s+offsetX)%%%.2f+offsetY", m.KeyExpressionMap()[UnknownKey], m.ModularUnit)
}

func (m ModularFunction) Init() {
	m.BaseFunc.BaseInit(ModularFunctionType)
}

func (m ModularFunction) Call(f float64) (float64, error) {
	if m.ModularUnit == 0 {
		return 0, ZeroValueError.Param("ModularFunction.modular_unit")
	}
	x, err := m.Keys()[UnknownKey].Call(f)
	if err != nil {
		return 0, err
	}
	if math.IsNaN(x) {
		return 0, NanValueError.Param("x")
	}

	return math.Mod(x+m.OffsetX, m.ModularUnit) + m.OffsetY, nil
}

// function initiator

const (
	ModularFunctionType    = "ModularFunctionType"
	TimeSecondFunctionType = "TimeSecondFunction"
)

var (
	ModularFunctionInitiator FuncInitiator = func() BaseFuncInterface {
		return &ModularFunction{
			BaseFunc: BaseFunc{
				DocValue: `ModularFunction is a modular function, to get the value % specify value. And the ModularFunction 
provide some derived function to get time.<br>

<ul>
<li> TimeSecondFunction: Always return value in 0-59. </li>
</ul>
`,
			},
		}
	}
	TimeSecondFunctionInitiator FuncInitiator = func() BaseFuncInterface {
		return &ModularFunction{
			BaseFunc: BaseFunc{
				IsDerivedVar: &TrueP,
				DocValue: `TimeSecondFunction always return value in 0-59, and set offset to zero(offsetX, offsetY).
`,
			},
			ModularUnit: 60,
			OffsetX:     0,
			OffsetY:     0,
		}
	}
)
