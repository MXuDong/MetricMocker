package funcs

import (
	"fmt"
	"math/rand"
)

type RandomFunction struct {
	BaseFunc

	BaseExpression string
	Seed           int64 `key:"seed" default:"0" mean:"The random seed"`

	randItem *rand.Rand
}

func (r *RandomFunction) Expression() string {
	if len(r.BaseExpression) == 0 {
		xExpression := r.KeyExpressionMap()[UnknownKey]
		r.BaseExpression = fmt.Sprintf("Random(%s)", xExpression)
	}

	return r.BaseExpression
}

func (r *RandomFunction) Init() {
	r.BaseInit(RandomFunctionType)
	r.randItem = rand.New(rand.NewSource(r.Seed))
}

func (r RandomFunction) Call(float642 float64) (float64, error) {
	return float64(r.randItem.Int()) + r.randItem.Float64(), nil
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
