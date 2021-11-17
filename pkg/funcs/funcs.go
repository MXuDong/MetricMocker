package funcs

// Function is the packing of the function like expression: 'y = x + 1'
// All the function is unary function.
//
// For input is type of int64, but return is float64, so it will lose some precision.
type Function interface {
	// Init the function
	Init(map[string]float64)

	// Execute return res of input x by function
	Execute(x int64) float64
}

// ResetAble mean the function can be reset,(some function need quick cal, need store some result, invoke it to clear)
type ResetAble interface {
	Reset()
}

const (
	Slope   = "slope"
	OffsetX = "offsetX"
	OffsetY = "offsetY"
)

// GetParam return the param value, if not exits, return default value
func GetParam(name string, param map[string]float64, defaultValue float64) float64 {
	if value, ok := param[name]; ok {
		return value
	}
	return defaultValue
}

// ZeroFunction always return zero
type ZeroFunction struct {
}

func (z ZeroFunction) Init(m map[string]float64) {
	// do nothing
}

func (z ZeroFunction) Execute(x int64) float64 {
	return 0
}
