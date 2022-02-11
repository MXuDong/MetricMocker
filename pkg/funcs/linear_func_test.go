package funcs

import (
	"fmt"
	"testing"
)

func TestBaseLinearFunction_Call(t *testing.T) {
	f := BaseLinearFunction{}
	f.Init(map[interface{}]interface{}{
		SlopeStr:   float32(10),
		OffsetXStr: float32(20),
		OffsetYStr: float32(30),
	})

	f.SetKeyFunc(UnknownKey, GenConstantValueFunc(1))
	fmt.Println(f.Expression())
	fmt.Println(f.Call(123))

}
