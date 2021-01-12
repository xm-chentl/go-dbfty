package utils

import "reflect"

// ToData 转换
func ToData(ds interface{}, screen func(item interface{}) interface{}) []interface{} {
	results := make([]interface{}, 0)
	dst := reflect.TypeOf(ds)
	dsv := reflect.ValueOf(ds)
	if dst.Kind() == reflect.Slice || dst.Kind() == reflect.Array {
		for i := 0; i < dsv.Len(); i++ {
			results = append(
				results,
				screen(dsv.Index(i).Interface()),
			)
		}
	}

	return results
}

// ForEach 循环
func ForEach(source []interface{}, iterator func(interface{}) interface{}) interface{} {
	return nil
}
