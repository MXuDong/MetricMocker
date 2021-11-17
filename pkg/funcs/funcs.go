package funcs

// Function is the packing of the function like expression: 'y = x + 1'
// All the function is unary function.
//
// For input is type of int64, but return is float64, so it will lose some precision.
type Function interface {

	// Execute return res of input x by function
	Execute(x int64) float64
}
