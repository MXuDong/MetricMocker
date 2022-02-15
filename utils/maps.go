package utils

import (
	"mmocker/utils/log"
	"reflect"
)

// IsMapSame return true when m1 == m2.
// If both of them is nil, return true.
// If one is nil, return false.
// When all the key and value is equals, return ture, else return false.
func IsMapSame(m1, m2 map[string]string) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil || m2 == nil {
		return false
	}

	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v1 != v2 {
			return false
		}
	}

	return true
}

// GetValueWithDefault return specify value from map. If specify value not found, return default value.
// But if specify value is empty value(can be found), it will return empty value.
func GetValueWithDefault(mapItem map[interface{}]interface{}, key interface{}, defaultValue interface{}) interface{} {
	if mapItem == nil {
		return defaultValue
	}

	if value, ok := mapItem[key]; ok {
		return value
	}
	return defaultValue
}

// GetStringWithDefault return string value from mapItem. If value not found, or can't convert to string. Return
// defaultValue.
func GetStringWithDefault(mapItem map[interface{}]interface{}, key interface{}, defaultValue string) string {
	value := GetValueWithDefault(mapItem, key, defaultValue)
	if data, ok := value.(string); ok {
		return data
	} else {
		log.Logger.Warnf("Convert to string error: %v.(%v) -> string", value, reflect.TypeOf(value))
	}
	return defaultValue
}

// GetFloat32WithDefault return float32 value from mapItem. If value not found, or can't convert to float32. Return
// defaultValue.
func GetFloat32WithDefault(mapItem map[interface{}]interface{}, key interface{}, defaultValue float32) (res float32) {
	value := GetValueWithDefault(mapItem, key, defaultValue)
	if data, ok := value.(float32); ok {
		return data
	}
	return defaultValue
}


func GetFloat64WithDefault(mapItem map[interface{}]interface{}, key interface{}, defaultValue float64)(res float64){
	value := GetValueWithDefault(mapItem, key, defaultValue)
	if data, ok := value.(float64); ok {
		return data
	}
	return defaultValue
}