package funcs

import (
	"math"
	"math/rand"
)

type RandomFunction struct {
	base     float64
	randSeed float64
	rang     float64
	ran      *rand.Rand

	params map[string]float64
}

func StandardRandomFunction(m map[string]float64) Function {
	r := RandomFunction{}
	r.Init(m)
	return &r
}

func (r *RandomFunction) Init(m map[string]float64) {
	if m == nil {
		r.base = 0
		r.randSeed = 0
		r.rang = 2
	} else {
		r.base = m[BasePoint]
		r.randSeed = m[Seed]
		r.rang = m[Range]
	}

	if r.rang < 1 {
		r.rang = 1
	}

	r.ran = rand.New(rand.NewSource(int64(r.randSeed)))
}

func (r *RandomFunction) Execute(x float64) float64 {
	return r.base + float64(r.ran.Int63n(int64(math.Ceil(r.rang)))) + r.ran.Float64()
}

func (r *RandomFunction) Params() map[string]float64 {
	return r.params
}
