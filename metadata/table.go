package metadata

import (
	"fmt"
	"reflect"
	"strings"
)

type table struct {
	typeOfData  reflect.Type
	valueOfDate reflect.Value
	name        string
	columns     []IColumn
	mapColumns  map[string]IColumn
	pkColumn    IColumn
	fkColumn    IColumn
}

func (t *table) Name() string {
	if t.name != "" {
		return t.name
	}
	// 反射是否存在自定义表名
	if _, ok := t.typeOfData.MethodByName(DEFAULTMETHOD); ok {
		results := t.valueOfDate.MethodByName(DEFAULTMETHOD).Call(nil)
		t.name = results[0].String()
	}
	if t.name == "" {
		t.name = strings.ToLower(t.typeOfData.Name())
	}

	return t.name
}

func (t table) GetType() reflect.Type {
	return t.typeOfData
}

func (t *table) GetColumns() []IColumn {
	if t.columns != nil {
		return t.columns
	}

	t.columns = make([]IColumn, 0)
	for index := 0; index < t.typeOfData.NumField(); index++ {
		t.columns = append(
			t.columns,
			&column{
				structOfField: t.typeOfData.Field(index),
			},
		)
	}

	return t.columns
}

func (t *table) GetColumnsByMap() map[string]IColumn {
	if t.mapColumns != nil {
		return t.mapColumns
	}
	t.mapColumns = make(map[string]IColumn)
	for _, column := range t.GetColumns() {
		t.mapColumns[column.Name()] = column
	}

	return t.mapColumns
}

func (t *table) GetPrimaryKeyBy() IColumn {
	if t.pkColumn != nil {
		return t.pkColumn
	}
	// 获取主键并缓存
	for _, column := range t.GetColumns() {
		if column.IsPrimaryKey() {
			t.fkColumn = column
			break
		}
	}

	return t.fkColumn
}

func (t table) Check() error {
	// 1. 是否存在主键
	if pk := t.GetPrimaryKeyBy(); pk == nil {
		return fmt.Errorf("The primary key does not exist")
	}

	return nil
}
