package funcs

import (
	"math"
	"time"
)

type LinearRangeTimePeak struct {
	offsetY float64
	offsetX float64
	rang    float64
	ratio   float64
}

func (l *LinearRangeTimePeak) Init(m map[string]float64) {
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

func (l *LinearRangeTimePeak) Execute(x float64) float64 {
	return l.ratio*(math.Mod(x+l.offsetX, l.rang)) + l.offsetY
}

func (l *LinearRangeTimePeak) Params() map[string]float64 {
	return map[string]float64{
		OffsetY: l.offsetY,
		OffsetX: l.offsetX,
		Ratio:   l.ratio,
		Range:   l.rang,
	}
}

func StandardLinearRangeTimePeak(m map[string]float64) Function {
	l := LinearRangeTimePeak{}
	l.Init(m)
	return &l
}

func RangeSecondLinearPeak() Function {
	l := LinearRangeTimePeak{}
	l.Init(map[string]float64{
		Range: 1 * time.Second.Seconds(),
		Ratio: 1,
	})
	return &l
}

func RangeMinuteLinearPeak() Function {
	l := LinearRangeTimePeak{}
	l.Init(map[string]float64{
		Range: 1 * time.Minute.Seconds(),
		Ratio: 1,
	})
	return &l
}

func RangeHourLinearPeak() Function {
	l := LinearRangeTimePeak{}
	l.Init(nil)
	return &l
}

func RangeDayLinearPeak() Function {
	l := LinearRangeTimePeak{}
	l.Init(map[string]float64{
		Range: 24 * time.Hour.Seconds(),
		Ratio: 1,
	})
	return &l
}
