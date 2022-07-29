package funcs

import (
	"fmt"
	"math"
)

const (
	SinFunctionType = "SinFunction"
)

// SinFunction y = Slope*SIN(x + OffsetX) +OffsetY
type SinFunction struct {
	BaseFunc
	BaseExpression string `expression:"y=Slope*SIN(x+OffsetX)+OffsetY"`

	Slope   float64 `mean:"Slope of Sin result" key:"slope" default:"1"`
	SlopeX  float64 `mean:"Slope of x" key:"slopeX" default:"1"`
	OffsetX float64 `mean:"offset of x" key:"offsetX" default:"0"`
	OffsetY float64 `mean:"offset of y" key:"offsetY" default:"0"`
}

func (s *SinFunction) Expression() string {
	if len(s.BaseExpression) == 0 {
		keyExpressionMap := map[string]string{}
		if s.HasKeyFunctions() {
			keyExpressionMap = s.BaseFunc.KeyExpressionMap()
		}
		s.BaseExpression = fmt.Sprintf("%v*SIN(%v*%v+%v)+%v", s.Slope, keyExpressionMap[UnknownKey], s.SlopeX, s.OffsetX, s.OffsetY)
	}
	return s.BaseExpression
}

func (s *SinFunction) Init() {
	s.BaseFunc.BaseInit(SinFunctionType)
}

func (s SinFunction) Call(f float64) (float64, error) {
	x, err := s.Keys()[UnknownKey].Call(f)
	if err != nil {
		return 0, err
	}
	return s.Slope*math.Sin(s.SlopeX*x+s.OffsetX) + s.OffsetY, nil
}

var (
	SinFunctionInitiator = func() BaseFuncInterface {
		return &SinFunction{
			BaseFunc: BaseFunc{
				DocValue: "SinFunction is a trigonometric function.",
			},
		}
	}
)