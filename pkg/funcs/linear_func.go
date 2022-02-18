package funcs

import (
	"fmt"
)

const (
	BaseLinearFunctionType = "BaseLinearFunction"
)

// BaseLinearFunction y = Slope * (x + OffsetX) + OffsetY
type BaseLinearFunction struct {
	BaseFunc
	BaseExpression string `expression:"y=Slope*(x+OffsetX)+OffsetY"`

	Slope   float64 `mean:"slope of line" key:"slope" default:"1"`
	OffsetX float64 `mean:"offset of x" key:"offsetX" default:"0"`
	OffsetY float64 `mean:"offset of y" key:"offsetY" default:"0"`
}

func (b BaseLinearFunction) Doc() string {
	return `
BaseLinearFunction is a simple one dimensional function.
`
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
	b.BaseFunc.BaseInit()
}

func (b BaseLinearFunction) Call(f float64) (float64, error) {
	x, err := b.Keys()[UnknownKey].Call(f)
	if err != nil {
		return 0, err
	}

	return b.Slope*(b.OffsetX+x) + b.OffsetY, nil
}

func (b BaseLinearFunction) Type() TypeStr {
	return BaseLinearFunctionType
}
