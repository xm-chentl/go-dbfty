package metadata

import "reflect"

// IColumn 字段接口
type IColumn interface {
	Name() string
	Field() string
	GetStruct() reflect.StructField
	GetValue(interface{}) (interface{}, bool)
	IsPrimaryKey() bool
	IsForeignKey() bool
	IsAuto() bool
}
