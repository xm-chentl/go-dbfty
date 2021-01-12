package metadata

import (
	"fmt"
	"reflect"
	"strings"
)

type column struct {
	name          string
	structOfField reflect.StructField
}

func (c *column) Name() string {
	if c.name != "" {
		return c.name
	}
	if c.structOfField.Tag != "" {
		c.name = c.structOfField.Tag.Get(TAGCOLUMN)
	}
	if c.name == "" {
		c.name = strings.ToLower(c.structOfField.Name)
	}

	return c.name
}

func (c column) Field() string {
	return fmt.Sprintf("`%s`", c.Name())
}

func (c column) GetStruct() reflect.StructField {
	return c.structOfField
}

func (c column) GetValue(data interface{}) (interface{}, bool) {
	valueOfData := reflect.ValueOf(data)
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		valueOfData = valueOfData.Elem()
	}

	// todo...临时处理的过滤器
	isNil := false
	rv := valueOfData.FieldByName(c.GetStruct().Name)
	switch rv.Kind() {
	case reflect.String:
		isNil = rv.String() != ""
		break
	case reflect.Interface:
		isNil = rv.Interface() != nil
		break
	case reflect.Int:
		isNil = rv.Int() != 0
		break
	}

	return rv.Interface(), isNil
}

func (c column) IsPrimaryKey() bool {
	_, ok := c.structOfField.Tag.Lookup(TAGPRIMARYKEY)
	return ok
}

func (c column) IsForeignKey() bool {
	_, ok := c.structOfField.Tag.Lookup(TAGFOREIGNKEY)
	return ok
}

func (c column) IsAuto() bool {
	_, ok := c.structOfField.Tag.Lookup(TAGAUTO)
	return ok
}
