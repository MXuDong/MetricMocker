package funcs

import (
	"fmt"
	"strings"
)

type BaseDivisionFunction struct {
	BaseFunc

	baseExpression string `expression:"y=(x+offsetX)/divisor+offsetY"`

	OffsetX float64 `key:"offsetX" mean:"offset of x." default:"0"`
	OffsetY float64 `key:"offsetY" mean:"offset of y." default:"0"`
	Divisor float64 `key:"divisor" mean:"divisor of function, can't be zero'" default:"1"`
}

func (b BaseDivisionFunction) Expression() string {
	// (x + offsetX)/divisor + offsetY
	expressionBytes := strings.Builder{}
	expressionBytes.WriteString("(")
	expressionBytes.WriteString(b.KeyExpressionMap()[UnknownKey])
	expressionBytes.WriteString(ExpressionOfValueWithSymbol(b.OffsetX))
	expressionBytes.WriteString(")/")
	expressionBytes.WriteString(fmt.Sprintf("%.2f", b.Divisor))
	expressionBytes.WriteString(ExpressionOfValueWithSymbol(b.OffsetY))

	return expressionBytes.String()
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

	return (x+b.OffsetX)/b.Divisor + b.OffsetY, nil
}

// function initiator

const (
	BaseDivisionFunctionType = "BaseDivisionFunction"
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
)