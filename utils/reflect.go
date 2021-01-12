package utils

import "reflect"

// GetTypeBySlice 获取切片的元素类型
func GetTypeBySlice(slice interface{}) reflect.Type {
	typeOfSlice := reflect.TypeOf(slice)
	if typeOfSlice.Kind() == reflect.Ptr {
		typeOfSlice = typeOfSlice.Elem()
	}

	return reflect.MakeSlice(typeOfSlice, 1, 1).Index(0).Type()
}
