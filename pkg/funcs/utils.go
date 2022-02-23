package funcs

import (
	"fmt"
	"reflect"
	"strconv"
)

func ExpressionOfValueWithSymbol(value float64) string {
	if value > 0 {
		return "+" + fmt.Sprintf("%f", value)
	} else if value == 0 {
		return ""
	} else {
		return "-" + fmt.Sprintf("%f", value)
	}
}

func ConvertMapStringToMapInterface(value map[string]string, keys map[string]FieldDescribe) (map[string]interface{}, error) {
	if value == nil || keys == nil {
		return map[string]interface{}{}, nil
	}
	if len(value) == 0 || len(keys) == 0 {
		return map[string]interface{}{}, nil
	}

	res := map[string]interface{}{}

	for k1, v1 := range value {
		if fieldDescribe, ok := keys[k1]; ok {
			if len(v1) == 0 {
				v1 = fieldDescribe.Default
			}
			switch fieldDescribe.Type {
			case reflect.Int64.String():
				if v, err := strconv.ParseInt(v1, 10, 64); err != nil {
					return res, err
				} else {
					res[k1] = v
				}
			case reflect.Bool.String():
				if v, err := strconv.ParseBool(v1); err != nil {
					return res, err
				} else {
					res[k1] = v
				}
			case reflect.Float64.String():
				if v, err := strconv.ParseFloat(v1, 64); err != nil {
					return res, err
				} else {
					res[k1] = v
				}
			case reflect.String.String():
				res[k1] = v1
			}
		}
	}

	return res, nil
}
