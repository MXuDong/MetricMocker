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

	params map[string]interface{}

	Slope   float64 `mean:"slope of line" key:"slope"`
	OffsetX float64 `mean:"offset of x" key:"offsetX"`
	OffsetY float64 `mean:"offset of y" key:"offsetY"`
}

func (b BaseLinearFunction) Doc() string {
	return `
BaseLinearFunction is a simple one dimensional function.
`
}

func (b BaseLinearFunction) KeyMap() map[string]struct{} {
	return map[string]struct{}{
		UnknownKey: {},
	}
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
	b.SetKeyFunc(UnknownKey, MetadataUnitFunction{})
}

func (b *BaseLinearFunction) Params() map[string]interface{} {
	if b.params == nil {
		b.params = GetParamMap(b)
	}
	return b.params
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
