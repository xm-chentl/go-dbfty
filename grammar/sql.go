package grammar

import (
	"fmt"
	"reflect"

	"github.com/xm-chentl/go-dbfty/metadata"
)

// Add 生成Add相关参数
func Add(data interface{}, fields []string, f func(table string, fields []string) string) (string, []interface{}, error) {
	table, err := metadata.Get(data)
	if err != nil {
		return "", nil, err
	}

	if fields == nil {
		fields = make([]string, 0)
		for _, column := range table.GetColumns() {
			fields = append(fields, column.Name())
		}
	}

	args := make([]interface{}, 0)
	valueOfData := reflect.ValueOf(data)
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		valueOfData = valueOfData.Elem()
	}
	for _, field := range fields {
		if column, ok := table.GetColumnsByMap()[field]; ok {
			args = append(
				args,
				valueOfData.FieldByName(column.GetStruct().Name).Interface(),
			)
		}
	}

	sql := f(table.Name(), fields)
	if err != nil {
		return "", nil, err
	}

	return sql, args, nil
}

// Delete 生成delete相关参数
func Delete(data interface{}, f func(table string) (string, []interface{})) (string, []interface{}, error) {
	table, err := metadata.Get(data)
	if err != nil {
		return "", nil, err
	}

	valueOfData := reflect.ValueOf(data)
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		valueOfData = valueOfData.Elem()
	}
	sql, args := f(table.Name())
	if args == nil {
		pkColumn := table.GetPrimaryKeyBy()
		sql += fmt.Sprintf(" `%s` = ?", pkColumn.Name())
		args = []interface{}{
			valueOfData.FieldByName(pkColumn.GetStruct().Name).Interface(),
		}
	}

	return sql, args, nil
}

// Update 生成update相关参数
func Update(data interface{}, sets []string, f func(table string, sets []string) (string, []interface{})) (string, []interface{}, error) {
	table, err := metadata.Get(data)
	if err != nil {
		return "", nil, err
	}

	if sets == nil {
		sets = make([]string, 0)
		for _, column := range table.GetColumns() {
			if !column.IsPrimaryKey() {
				sets = append(sets, column.Name())
			}
		}
	}

	args := make([]interface{}, 0)
	valueOfData := reflect.ValueOf(data)
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		valueOfData = valueOfData.Elem()
	}
	for _, field := range sets {
		if column, ok := table.GetColumnsByMap()[field]; ok {
			args = append(
				args,
				valueOfData.FieldByName(column.GetStruct().Name).Interface(),
			)
		}
	}

	sql, argsOfWhere := f(table.Name(), sets)
	if len(argsOfWhere) == 0 {
		sql += fmt.Sprintf("`%s` = ?", table.GetPrimaryKeyBy().Name())
		args = append(
			args,
			valueOfData.FieldByName(table.GetPrimaryKeyBy().GetStruct().Name).Interface(),
		)
	} else {
		args = append(args, argsOfWhere...)
	}

	return sql, args, nil
}

// Query 查询
func Query(data interface{}, f func(table string) (string, []interface{})) (string, []interface{}, error) {
	// SELECT 字段/* FROM 表名 WHERE
	table, err := metadata.Get(data)
	if err != nil {
		return "", nil, err
	}

	sql, args := f(table.Name())

	return sql, args, nil
}
