package funcs

import (
	"fmt"
	"mmocker/utils"
)

const (
	BaseLinearFunctionType = "BaseLinearFunction"
)

// BaseLinearFunction y = slope * (x + offsetX) + offsetY
type BaseLinearFunction struct {
	BaseFunc
	baseExpression string

	paramsMap map[string]interface{}

	slope   float64
	offsetX float64
	offsetY float64
}

func (b BaseLinearFunction) KeyMap()map[string]struct{}{
	return map[string]struct{}{
		UnknownKey: {},
	}
}

func (b BaseLinearFunction) Expression() string {
	keyExpressionMap := map[string]string{}
	if b.HasKeyFunctions() {
		keyExpressionMap = b.BaseFunc.KeyExpressionMap()
	}
	return fmt.Sprintf("%v*(%v+%v)+%v", b.paramsMap[SlopeStr], keyExpressionMap[UnknownKey], b.paramsMap[OffsetXStr], b.paramsMap[OffsetYStr])
}

func (b *BaseLinearFunction) Init(m map[interface{}]interface{}) {
	b.slope = utils.GetFloat64WithDefault(m, SlopeStr, 1)
	b.offsetX = utils.GetFloat64WithDefault(m, OffsetXStr, 0)
	b.offsetY = utils.GetFloat64WithDefault(m, OffsetYStr, 0)
	b.paramsMap = map[string]interface{}{
		SlopeStr:   b.slope,
		OffsetXStr: b.offsetX,
		OffsetYStr: b.offsetY,
	}

	b.SetKeyFunc(UnknownKey, MetadataUnitFunction{})
}

func (b BaseLinearFunction) Params() map[string]interface{} {
	return b.paramsMap
}

func (b BaseLinearFunction) Call(f float64) (float64, error) {
	x, err := b.Keys()[UnknownKey].Call(f)
	if err != nil {
		return 0, err
	}

	return b.slope*(b.offsetX+x) + b.offsetY, nil
}

func (b BaseLinearFunction) Type() TypeStr {
	return BaseLinearFunctionType
}
