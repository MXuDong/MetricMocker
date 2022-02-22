package funcs

import (
	"fmt"
	"math"
	"strings"
)

type BaseDivisionFunction struct {
	BaseFunc

	baseExpression string  `expression:"y=(x+offsetX)/divisor+offsetY"`
	Divisor        float64 `key:"divisor" mean:"divisor of function, can't be zero'" default:"1"`
}

func (b *BaseDivisionFunction) Expression() string {
	if len(b.baseExpression) == 0 {
		// (x + offsetX)/divisor + offsetY
		expressionBytes := strings.Builder{}
		expressionBytes.WriteString(b.KeyExpressionMap()[UnknownKey])
		expressionBytes.WriteString("/")
		expressionBytes.WriteString(fmt.Sprintf("%.2f", b.Divisor))
		b.baseExpression = expressionBytes.String()
	}
	return b.baseExpression
}

func (b *BaseDivisionFunction) Init() {
	b.BaseInit(BaseDivisionFunctionType)
}

func (b BaseDivisionFunction) Call(float642 float64) (float64, error) {
	if b.Divisor == 0 {
		return 0, ZeroValueError.Param("divisor")
	}

	x, err := b.Keys()[UnknownKey].Call(float642)
	if err != nil {
		return 0, err
	}

	return x / b.Divisor, nil
}

type FloatFloorFunction struct {
	BaseFunc
}

func (f FloatFloorFunction) Expression() string {
	xExpression := f.KeyExpressionMap()[UnknownKey]
	return fmt.Sprintf("(%s)(Floor)", xExpression)
}

func (f *FloatFloorFunction) Init() {
	f.BaseInit(FloatFloorFunctionType)
}

func (f FloatFloorFunction) Call(float642 float64) (float64, error) {
	x, err := f.Keys()[UnknownKey].Call(float642)
	if err != nil {
		return 0, err
	}
	return math.Floor(x), nil
}

// function initiator

const (
	BaseDivisionFunctionType = "BaseDivisionFunction"
	FloatFloorFunctionType   = "FloatFloorFunction"
)

var (
	BaseDivisionFunctionInitiator FuncInitiator = func() BaseFuncInterface {
		return &BaseDivisionFunction{
			BaseFunc: BaseFunc{
				DocValue: `The base division function is basic function. And the divisor can't be zero.<br>
It will return error when x / 0.`,
			},
		}
	}
	FloatFloorFunctionInitiator FuncInitiator = func() BaseFuncInterface {
		return &FloatFloorFunction{
			BaseFunc{
				DocValue: `FloatFloorFunction returns the greatest integer value less than or equal to x. `,
			},
		}
	}
)
