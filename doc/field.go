package doc
//
//import (
//	"mmocker/pkg/funcs"
//	"reflect"
//)
//
//const (
//	ExpressionTag = "expression"
//	MinValueTag   = "minValue"
//)
//
//func GetExpression(funcInterface funcs.BaseFuncInterface) string {
//	if funcInterface == nil {
//		return ""
//	}
//
//	typ := reflect.TypeOf(funcInterface)
//
//	for index := 0; index < typ.NumField(); index++ {
//		fieldItem := typ.Field(index)
//		fieldItem.Tag.Lookup()
//	}
//}
