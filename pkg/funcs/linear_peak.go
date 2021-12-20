package funcs

import (
	"math"
	"time"
)

type LinearPeak struct {
	offsetY float64
	offsetX float64
	rang    float64
	ratio   float64
}

func (l *LinearPeak) Init(m map[string]float64) {
	if m == nil {
		l.offsetY = 0
		l.rang = 1 * time.Hour.Seconds()
		l.offsetX = 0
		l.ratio = 1
	} else {
		l.offsetY = m[OffsetY]
		l.rang = m[Range]
		l.offsetX = m[OffsetX]
		l.ratio = m[Ratio]
	}
}

func (l *LinearPeak) Execute(x float64) float64 {
	return l.ratio*(math.Mod(x+l.offsetX, l.rang)) + l.offsetY
}

func (l *LinearPeak) Params() map[string]float64 {
	return map[string]float64{
		OffsetY: l.offsetY,
		OffsetX: l.offsetX,
		Ratio:   l.ratio,
		Range:   l.rang,
	}
}

func StandardLinearPeak(m map[string]float64) Function {
	l := LinearPeak{}
	l.Init(m)
	return &l
}

func SecondLinearPeak() Function {
	l := LinearPeak{}
	l.Init(map[string]float64{
		Range: 1 * time.Second.Seconds(),
		Ratio: 1,
	})
	return &l
}

func MinuteLinearPeak() Function {
	l := LinearPeak{}
	l.Init(map[string]float64{
		Range: 1 * time.Minute.Seconds(),
		Ratio: 1,
	})
	return &l
}

func HourLinearPeak() Function {
	l := LinearPeak{}
	l.Init(nil)
	return &l
}

func DayLinearPeak() Function {
	l := LinearPeak{}
	l.Init(map[string]float64{
		Range: 24 * time.Hour.Seconds(),
		Ratio: 1,
	})
	return &l
}
