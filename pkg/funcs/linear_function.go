package funcs

// LinearFunction is the simple linear-function.
// the func is : y = slope * x + offset
type LinearFunction struct {
	slope  float64
	offset float64
}

func (l *LinearFunction) Execute(x int64) float64 {
	return l.slope*float64(x) + l.offset
}

// NewLinearFunction return new linear-function of special slope and offset.
func NewLinearFunction(slope, offset float64) Function {
	return &LinearFunction{
		slope:  slope,
		offset: offset,
	}
}

// DefaultLinearFunction return function with expression is : y = x
func DefaultLinearFunction() Function {
	return NewLinearFunction(1, 0)
}

// ReverseLinearFunction return function with expression is : y = -x
func ReverseLinearFunction() Function {
	return NewLinearFunction(-1, 0)
}

func (l *LinearFunction) Slope() float64 {
	return l.slope
}

func (l *LinearFunction) Offset() float64 {
	return l.offset
}
