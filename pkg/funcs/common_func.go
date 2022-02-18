package funcs

import (
	"fmt"
	"strconv"
)

const (
	StartZeroFuncType = "StartZeroFunc"
)

// ==================== start zero func

// StartZeroFunc used to reset x by time.
type StartZeroFunc struct {
	BaseFunc

	StartTime float64 `key:"start_value" mean:"Where value will start" default:"-1.0"`
}

func (s StartZeroFunc) Doc() string {
	return `
StartZeroFuncType always start with 0. And start work when first call. Used to reset UnknownKey(x) by time.
`
}

func (s StartZeroFunc) Expression() string {
	return fmt.Sprintf("(%s->start_value)", s.KeyExpressionMap()[UnknownKey])
}

func (s *StartZeroFunc) Init() {
	// do nothing
	s.StartTime = -1
	s.BaseFunc.BaseInit(StartZeroFuncType)
}

func (s StartZeroFunc) Params() map[string]interface{} {
	return map[string]interface{}{}
}

func (s *StartZeroFunc) Call(f float64) (float64, error) {
	x, err := s.Keys()[UnknownKey].Call(f)
	if err != nil {
		return 0, err
	}
	if s.StartTime == -1 {
		s.StartTime = x
	}
	return x - s.StartTime, nil
}

// ==================== metadata unit function

const MetadataUnitFunctionType = "MetadataUnitFunction"

// MetadataUnitFunction is the base function of key
type MetadataUnitFunction struct {
}

func (m MetadataUnitFunction) SetType(str TypeStr){
	// do nothing
}

// IsDerived always return false, because the MetadataUnitFunction has no derived function.
func (m MetadataUnitFunction) IsDerived() bool {
	return false
}

func (m MetadataUnitFunction) Doc() string {
	return `
MetadataUnitFunction is a base metadata unit function. Use to define the UnknownKey(x).
`
}

func (m MetadataUnitFunction) KeyMap() map[string]struct{} {
	return map[string]struct{}{}
}

func (m MetadataUnitFunction) Type() TypeStr {
	return MetadataUnitFunctionType
}

func (m MetadataUnitFunction) Expression() string {
	return UnknownKey
}

func (m MetadataUnitFunction) Init() {

	// the MetadataUnitFunction has no params
}
func (m *MetadataUnitFunction) InitP() BaseFuncInterface {
	m.Init()
	return m
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

func (m MetadataUnitFunction) Call(f float64) (float64, error) {
	return f, nil
}

// ==================== constant value function

func GenConstantValueFunc(value float64) BaseFuncInterface {
	return InitFunction(&ConstantValueFunction{}, map[string]interface{}{"Value": value})
}

// ConstantValueFunction always return a constant value from init-func.
type ConstantValueFunction struct {
	BaseFunc
	BaseExpression string  `expression:"x(=value)"`
	Value          float64 `key:"Value"`
}

func (c ConstantValueFunction) Doc() string {
	return `
ConstantValueFunction use to return constant value. And the expression is keep one value.`
}

func (b ConstantValueFunction) KeyMap() map[string]struct{} {
	return map[string]struct{}{}
}

func (b ConstantValueFunction) Type() TypeStr {
	return "ConstantValueFunction"
}

// Expression return value directly.
func (b ConstantValueFunction) Expression() string {
	return strconv.FormatFloat(b.Value, 'f', -1, 64)
}

// Init a ConstantValueFunction with value, it can be call more than one time. It will change value.
// Need param: ValueStr(float64)
func (b *ConstantValueFunction) Init() {
}

// Params return value directly.
func (b ConstantValueFunction) Params() map[string]interface{} {
	return map[string]interface{}{
		ValueStr: b.Value,
	}
}

func (b ConstantValueFunction) Keys() map[string]BaseFuncInterface {
	return map[string]BaseFuncInterface{}
}

// Call return value directly
func (b ConstantValueFunction) Call(f float64) (float64, error) {
	return b.Value, nil
}
