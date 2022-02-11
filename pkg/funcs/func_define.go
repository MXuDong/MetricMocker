package funcs

import (
	"mmocker/utils"
	"strconv"
)

// BaseFuncInterface define the func. Is the abstract of the instance
type BaseFuncInterface interface {

	// Expression return the expression of function. And all the params should with value, not the symbol.
	// For example:
	// 		If the function is : y = a(x+b) + c, the a, b & c is the symbol. And the params is a=2, b=3 & c=4
	//		The return expression: y=2*(x+3)+4 or y=2(x+3)+4
	//		Not the : y = a(x+b) +c or y=2*x+10.
	//		Because the b is the offset for x, and the c is offset for y.
	// And if func set key function. The expression should return expression replace the key with key-function.
	// For example：
	// 		If the function is : y = a1(x1) + a2(x2), the a1 & a2 is the symbol. The x1 & x2 is the key.
	//		And, set a1=2, a2=3, x1 with key function: 5*x3, x2 with key function: 6+x4
	//		The return expression: y=2*(5*x)+3*(6+x)
	Expression() string

	// Init set the params of function, and all the function can set default value if some symbol not provide.
	Init(map[interface{}]interface{})

	// SetKeyFunc will set the key function. All the key func will calculate first. The sample level of key-function
	//will be seen as same key. If expression is y=x1+x2, the x1 and x2 is different key. But when set key-function
	//with x1(to k1) and x2(to k2), the expression is: y=k1(x)+k2(x).
	// If bFunc is nil, it will delete specify key function.
	SetKeyFunc(key string, bFunc BaseFuncInterface)

	// Params return the params of function. If init not support some symbol value, it will return default value.
	Params() map[string]interface{}

	// Keys return the keys of function. If the function already bind with key-function, the res contain the key from
	//key-functions.
	Keys() map[string]BaseFuncInterface

	// Call return the res of specify key value. But if function has more keys, it will set same key to call. If
	//function want call by different key value, use CallWithValue to get res.
	Call(float32) (float32, error)
}

type BaseFunc struct {
	keyFunctions map[string]BaseFuncInterface
}

func (bf *BaseFunc) SetConstantValue(key string, value float32) {
	bf.SetKeyFunc(key, (&ConstantValueFunction{}).InitP(map[interface{}]interface{}{ValueStr: value}))
}

func (bf *BaseFunc) SetKeyFunc(key string, bFunc BaseFuncInterface) {
	if bf.keyFunctions == nil {
		bf.keyFunctions = map[string]BaseFuncInterface{}
	}
	if bFunc == nil {
		if _, ok := bf.keyFunctions[key]; ok {
			delete(bf.keyFunctions, key)
		}
		return
	}
	bf.keyFunctions[key] = bFunc
}
func (bf BaseFunc) Keys() map[string]BaseFuncInterface {
	return bf.keyFunctions
}

func (bf *BaseFunc) DoFuncWithDefaultValue(key string, param float32, defaultValue float32) (float32, error) {
	if funcItem, ok := bf.keyFunctions[key]; ok {
		return funcItem.Call(param)
	}
	return defaultValue, nil
}

// HasKeyFunctions return true when BaseFunc has key-functions.
func (bf *BaseFunc) HasKeyFunctions() bool {
	return len(bf.keyFunctions) != 0
}

// KeyExpressionMap return the map of keys.
func (bf *BaseFunc) KeyExpressionMap() map[string]string {
	res := map[string]string{}
	if !bf.HasKeyFunctions() {
		return res
	}
	for key, funcItem := range bf.keyFunctions {
		res[key] = (funcItem).Expression()
	}

	return res
}

// ======== constant value function

func GenConstantValueFunc(value float32) BaseFuncInterface {
	return (&ConstantValueFunction{}).InitP(map[interface{}]interface{}{ValueStr: value})
}

// ConstantValueFunction always return a constant value from init-func.
type ConstantValueFunction struct {
	BaseFunc
	value float32
}

// Expression return value directly.
func (b ConstantValueFunction) Expression() string {
	return strconv.FormatFloat(float64(b.value), 'f', -1, 32)
}

// Init a ConstantValueFunction with value, it can be call more than one time. It will change value.
// Need param: ValueStr(float32)
func (b *ConstantValueFunction) Init(m map[interface{}]interface{}) {
	b.value = utils.GetFloat32WithDefault(m, ValueStr, 0)
}

func (b *ConstantValueFunction) InitP(m map[interface{}]interface{}) BaseFuncInterface {
	b.Init(m)
	return b
}

// Params return value directly.
func (b ConstantValueFunction) Params() map[string]interface{} {
	return map[string]interface{}{
		ValueStr: b.value,
	}
}

func (b ConstantValueFunction) Keys() map[string]BaseFuncInterface {
	return map[string]BaseFuncInterface{}
}

// Call return value directly
func (b ConstantValueFunction) Call(f float32) (float32, error) {
	return b.value, nil
}

// ================ base metadata unit function

// MetadataUnitFunction is the base function of key
type MetadataUnitFunction struct {
}

func (m MetadataUnitFunction) Expression() string {
	return UnknownKey
}

func (m MetadataUnitFunction) Init(m2 map[interface{}]interface{}) {

	// the MetadataUnitFunction has no params
}

func (m MetadataUnitFunction) SetKeyFunc(key string, bFunc BaseFuncInterface) {
	// the metadata unit function has no key
}

func (m MetadataUnitFunction) Params() map[string]interface{} {
	return map[string]interface{}{}
}

func (m MetadataUnitFunction) Keys() map[string]BaseFuncInterface {
	return map[string]BaseFuncInterface{}
}

func (m MetadataUnitFunction) Call(f float32) (float32, error) {
	return f, nil
}

// ======== the same params
const (
	UnknownKey = "x"

	SlopeStr   = "slope"
	OffsetXStr = "offset_x"
	OffsetYStr = "offset_y"
	ValueStr   = "value"
)

// TODO: Complate
type CalculateError struct {
	Reason string
}

func (ce *CalculateError) Error() string {
	return ce.Reason
}
