package funcs

import (
	"fmt"
)

// BaseLinearFunction y = Slope * (x + OffsetX) + OffsetY
type BaseLinearFunction struct {
	BaseFunc
	BaseExpression string `expression:"y=Slope*(x+OffsetX)+OffsetY"`

	Slope   float64 `mean:"slope of line" key:"slope" default:"1"`
	OffsetX float64 `mean:"offset of x" key:"offsetX" default:"0"`
	OffsetY float64 `mean:"offset of y" key:"offsetY" default:"0"`
}

func (b *BaseLinearFunction) Expression() string {
	if len(b.BaseExpression) == 0 {
		keyExpressionMap := map[string]string{}
		if b.HasKeyFunctions() {
			keyExpressionMap = b.BaseFunc.KeyExpressionMap()
		}
		b.BaseExpression = fmt.Sprintf("%v*(%v+%v)+%v", b.Slope, keyExpressionMap[UnknownKey], b.OffsetX, b.OffsetY)
	}

	return b.BaseExpression
}

func (b *BaseLinearFunction) Init() {
	b.BaseFunc.BaseInit(BaseLinearFunctionType)
}

func (b BaseLinearFunction) Call(f float64) (float64, error) {
	x, err := b.Keys()[UnknownKey].Call(f)
	if err != nil {
		return 0, err
	}

	return b.Slope*(b.OffsetX+x) + b.OffsetY, nil
}

// ============ function initiators

const (
	BaseLinearFunctionType          = "BaseLinearFunction"
	SingleLinearFunctionType        = "SingleLinearFunction"
	ReverseSingleLinearFunctionType = "ReverseSingleLinearFunction"
)

var (
	BaseLinearFunctionInitiator = func() BaseFuncInterface {
		return &BaseLinearFunction{
			BaseFunc: BaseFunc{
				DocValue: `BaseLinearFunction is a simple one dimensional function. The default is type: 
SingleLinearFunction function.`,
			},
		}
	}
	SingleLinearFunctionInitiator = func() BaseFuncInterface {
		return &BaseLinearFunction{
			BaseFunc: BaseFunc{
				IsDerivedVar: &TrueP,
				DocValue: `BaseLinearFunction is a simple one dimensional function. The default is type: 
SingleLinearFunction function.`,
			},
			Slope:   1,
			OffsetX: 0,
			OffsetY: 0,
		}
	}
	ReverseSingleLinearFunctionInitiator = func() BaseFuncInterface {
		return &BaseLinearFunction{
			BaseFunc: BaseFunc{
				IsDerivedVar: &TrueP,
				DocValue: `BaseLinearFunction is a simple one dimensional function. The default is type: 
SingleLinearFunction function.`,
			},
			Slope:   -1,
			OffsetX: 0,
			OffsetY: 0,
		}
	}
)
