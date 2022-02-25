package funcs

import (
	"fmt"
	"math/rand"
)

type RandomFunction struct {
	BaseFunc

	BaseExpression string
	Seed           int64   `key:"seed" default:"0" mean:"The random seed."`
	Range          int64   `key:"range" default:"1" mean:"The range for random value."`
	OffsetY        float64 `key:"offsetY" default:"0" mean:"The offset of y."`

	randItem *rand.Rand
}

func (r *RandomFunction) Expression() string {
	if len(r.BaseExpression) == 0 {
		xExpression := r.KeyExpressionMap()[UnknownKey]
		r.BaseExpression = fmt.Sprintf("Random(%s)%%%d%s", xExpression, r.Range, ExpressionOfValueWithSymbol(r.OffsetY))
	}

	return r.BaseExpression
}

func (r *RandomFunction) Init() {
	r.BaseInit(RandomFunctionType)
	r.randItem = rand.New(rand.NewSource(r.Seed))
}

func (r RandomFunction) Call(float642 float64) (float64, error) {
	if r.Range == 0 {
		return 0 + r.OffsetY, nil
	}
	if r.Range < 0 {
		return 0, ShouldBiggerThanZero.Param("range")
	}
	value := float64(r.randItem.Int63n(r.Range))
	value += r.randItem.Float64()
	value += r.OffsetY
	return value, nil
}

const (
	RandomFunctionType = "RandomFunction"
)

var (
	RandomFunctionInitiator FuncInitiator = func() BaseFuncInterface {
		return &RandomFunction{
			BaseFunc: BaseFunc{
				DocValue: `Random function return random value, every call will return different value. But if seed is 
same value, call by same order will return same value.`,
			},
		}
	}
)
