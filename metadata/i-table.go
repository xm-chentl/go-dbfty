package metadata

import "reflect"

// ITable 表接口
type ITable interface {
	Name() string
	Check() error
	GetColumns() []IColumn
	GetColumnsByMap() map[string]IColumn
	GetType() reflect.Type
	GetPrimaryKeyBy() IColumn
	GetValueByMap(interface{}) map[string]interface{}
}
