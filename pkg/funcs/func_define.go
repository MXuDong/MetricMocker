package funcs

import (
	"mmocker/utils/log"
	"reflect"
	"strconv"
)

type TypeStr string

// BaseFuncInterface define the func. Is the abstract of the instance
type BaseFuncInterface interface {
	Type() TypeStr

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
	// The Expression should static, in other words, it should return same value if invoke more times.
	Expression() string

	// Init will be invoked later than set param. The init operator should not update param. Some time when function
	// need reInit should invoke it. In others words, invoke it mean function should re-start work when next call.
	Init()

	// Doc should return the function describe. Such like usage, params values. More times, it will show on html page.
	Doc() string

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

	KeyMap() map[string]struct{}

	// Call return the res of specify key value. But if function has more keys, it will set same key to call. If
	//function want call by different key value, use CallWithValue to get res.
	Call(float642 float64) (float64, error)
}

type BaseFunc struct {
	keyFunctions map[string]BaseFuncInterface
}

func (bf *BaseFunc) SetConstantValue(key string, value float64) {
	bf.SetKeyFunc(key, InitFunction(&ConstantValueFunction{}, map[string]interface{}{"Value": value}))
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

func (bf *BaseFunc) DoFuncWithDefaultValue(key string, param float64, defaultValue float64) (float64, error) {
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

// ================ base metadata unit function

// InitFunction will init a function, use reflect. In the end of IniFunction, it will invoke BaseFuncInterface.Init()
func InitFunction(funcItem BaseFuncInterface, params map[string]interface{}) BaseFuncInterface {
	if funcItem == nil {
		return MetadataUnitFunction{}
	}

	typ := reflect.TypeOf(funcItem).Elem()
	valu := reflect.ValueOf(funcItem).Elem()
	for fieldIndex := 0; fieldIndex < typ.NumField(); fieldIndex++ {
		fieldItem := typ.Field(fieldIndex)
		keyName, ok := GetParamKey(fieldItem)
		if !ok {
			continue
		}

		// set default value(is has ParamDefaultKey tag)
		if defaultValueStr, hasDefaultTag := fieldItem.Tag.Lookup(ParamDefaultKey); hasDefaultTag {
			switch valu.FieldByName(fieldItem.Name).Kind() {
			case reflect.Int64:
				if defaultValue, err := strconv.ParseInt(defaultValueStr, 10, 64); err == nil {
					valu.FieldByName(fieldItem.Name).SetInt(defaultValue)
				} else {
					log.Logger.Errorf("%v", err)
				}
			case reflect.Float64:
				if defaultValue, err := strconv.ParseFloat(defaultValueStr, 64); err == nil {
					valu.FieldByName(fieldItem.Name).SetFloat(defaultValue)
				} else {
					log.Logger.Errorf("%v", err)
				}
			case reflect.Bool:
				if defaultValue, err := strconv.ParseBool(defaultValueStr); err == nil {
					valu.FieldByName(fieldItem.Name).SetBool(defaultValue)
				} else {
					log.Logger.Errorf("%v", err)
				}
			case reflect.String:
				valu.FieldByName(fieldItem.Name).SetString(defaultValueStr)
			}
		}

		// parse from params
		if paramValue, containKey := params[keyName]; containKey {
			switch valu.FieldByName(fieldItem.Name).Kind() {
			case reflect.Int64:
				if v, convertFlag := paramValue.(int64); convertFlag {
					valu.FieldByName(fieldItem.Name).SetInt(v)
				}
			case reflect.Float64:
				if v, convertFlag := paramValue.(float64); convertFlag {
					valu.FieldByName(fieldItem.Name).SetFloat(v)
				}
			case reflect.Bool:
				if v, convertFlag := paramValue.(bool); convertFlag {
					valu.FieldByName(fieldItem.Name).SetBool(v)
				}
			case reflect.String:
				if v, convertFlag := paramValue.(string); convertFlag {
					valu.FieldByName(fieldItem.Name).SetString(v)
				}
			}
		}
	}

	// complete init, and first invoke init func.
	funcItem.Init()

	return funcItem
}

func GetParamMap(funcInterface BaseFuncInterface) map[string]interface{} {
	if funcInterface == nil {
		return map[string]interface{}{}
	}

	res := map[string]interface{}{}
	typ := reflect.TypeOf(funcInterface).Elem()
	valu := reflect.ValueOf(funcInterface).Elem()

	for fieldIndex := 0; fieldIndex < typ.NumField(); fieldIndex++ {
		fieldItem := typ.Field(fieldIndex)
		keyName, ok := GetParamKey(fieldItem)
		if !ok {
			continue
		}
		fieldOperator := valu.FieldByName(fieldItem.Name)
		switch fieldOperator.Kind() {
		case reflect.Int64:
			res[keyName] = fieldOperator.Int()
		case reflect.Float64:
			res[keyName] = fieldOperator.Float()
		case reflect.Bool:
			res[keyName] = fieldOperator.Bool()
		case reflect.String:
			res[keyName] = fieldOperator.String()
		}
	}
	return res
}

func GetParamKey(fs reflect.StructField) (string, bool) {
	if keyName, ok := fs.Tag.Lookup(ParamKeyKey); ok {
		if len(keyName) == 0 {
			return fs.Name, ok
		}
		return keyName, ok
	}
	return "", false
}

func GetParamFields(funcInterface BaseFuncInterface) map[string]FieldDescribe {
	if funcInterface == nil {
		return map[string]FieldDescribe{}
	}

	typ := reflect.TypeOf(funcInterface).Elem()
	res := map[string]FieldDescribe{}

	for fieldIndex := 0; fieldIndex < typ.NumField(); fieldIndex++ {
		fieldItem := typ.Field(fieldIndex)

		if keyName, ok := GetParamKey(fieldItem); ok {
			meanValue := fieldItem.Tag.Get(ParamMeanKey)
			defaultValue := fieldItem.Tag.Get(ParamDefaultKey)
			fieldType := fieldItem.Type.Name()
			fieldDescribe := FieldDescribe{
				KeyName: keyName,
				Mean:    meanValue,
				Default: defaultValue,
				Type:    fieldType,
			}
			res[keyName] = fieldDescribe
		}
	}

	return res
}

func GetExpression(funcInterface BaseFuncInterface) string {
	if funcInterface == nil {
		return ""
	}

	typ := reflect.TypeOf(funcInterface).Elem()
	expression := ""

	for fieldIndex := 0; fieldIndex < typ.NumField(); fieldIndex++ {
		fieldItem := typ.Field(fieldIndex)
		if value, ok := fieldItem.Tag.Lookup(ExpressionKey); ok {
			if len(expression) != 0 {
				expression += "|"
			}
			expression += value
		}
	}

	if len(expression) == 0 {
		expression = "y=" + funcInterface.Expression()
	}

	return expression
}

// ======== the same params
const (
	UnknownKey = "x"

	SlopeStr   = "slope"
	OffsetXStr = "offset_x"
	OffsetYStr = "offset_y"
	ValueStr   = "value"

	ParamKeyKey     = "key"  // all func param should have this tag, if empty of value, use filed name as key.
	ParamMeanKey    = "mean" // the param usage.
	ParamDefaultKey = "default"
	ExpressionKey   = "expression" // the function expression. If has no expression, invoke BaseFuncInterface.Expression()
)

// field properties

// FieldDescribe contain the key tags.
type FieldDescribe struct {
	KeyName string
	Mean    string
	Default string
	Type    string
}

// TODO: Complate
type CalculateError struct {
	Reason string
}

func (ce *CalculateError) Error() string {
	return ce.Reason
}
