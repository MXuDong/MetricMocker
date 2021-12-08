package funcs

// LinearFunction is the simple linear-function.
// the func is : y = slope * (x + offsetX) + offsetY
type LinearFunction struct {
	slope   float64
	offsetX float64
	offsetY float64
}

func (l *LinearFunction) Execute(x float64) float64 {
	return l.slope*(x+l.offsetX) + l.offsetY
}

// Init linear function, need slope, offsetX and offsetY
func (l *LinearFunction) Init(param map[string]float64) {
	l.slope = GetParam(Slope, param, 0)
	l.offsetX = GetParam(OffsetX, param, 0)
	l.offsetY = GetParam(OffsetY, param, 0)
}

func NewLinearFunctionByMap(param map[string]float64) Function {
	l := LinearFunction{}
	l.Init(param)
	return &l
}

// NewLinearFunction return new linear-function of special slope and offset.
func NewLinearFunction(slope, offsetX, offsetY float64) Function {
	l := LinearFunction{}
	initParam := map[string]float64{
		Slope:   slope,
		OffsetX: offsetX,
		OffsetY: offsetY,
	}
	l.Init(initParam)
	return &l
}
func (l *LinearFunction) Params() map[string]float64 {
	return map[string]float64{
		"slope":   l.Slope(),
		"offsetX": l.OffsetX(),
		"offsetY": l.OffsetY(),
	}
}

func (l *LinearFunction) Slope() float64 {
	return l.slope
}

func (l *LinearFunction) OffsetX() float64 {
	return l.offsetX
}

func (l *LinearFunction) OffsetY() float64 {
	return l.offsetY
}

// DefaultLinearFunction return function with expression is : y = x
func DefaultLinearFunction() Function {
	return NewLinearFunction(1, 0, 0)
}

// ReverseLinearFunction return function with expression is : y = -x
func ReverseLinearFunction() Function {
	return NewLinearFunction(-1, 0, 0)
}
