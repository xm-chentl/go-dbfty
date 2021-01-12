package metadata

import (
	"reflect"
	"testing"
)

func Test_Name_有Tag(t *testing.T) {
	field, ok := getTestPersonFieldStruct("Name")
	if !ok {
		t.Error("err")
	}
	column := column{
		structOfField: field,
	}
	if n := column.Name(); n != "name" {
		t.Error("err")
	}
	if column.Name() != column.name {
		t.Error("err")
	}
}

func Test_Name_无Tag(t *testing.T) {
	field, ok := getTestPersonFieldStruct("ID")
	if !ok {
		t.Error("err")
	}
	column := column{
		structOfField: field,
	}
	if n := column.Name(); n != "ID" {
		t.Error("err")
	}
}

func Test_GetStruct(t *testing.T) {
	field, ok := getTestPersonFieldStruct("Age")
	if !ok {
		t.Error("err")
	}
	column := column{
		structOfField: field,
	}
	ageOfStruct := column.GetStruct()
	if ageOfStruct.Name != "Age" {
		t.Error("err")
	}
	if ageOfStruct.Tag.Get(TAGCOLUMN) != "age" {
		t.Error("err")
	}
}

func Test_IsPrimaryKey_有(t *testing.T) {
	field, ok := getTestPersonFieldStruct("ID")
	if !ok {
		t.Error("err")
	}
	column := column{
		structOfField: field,
	}
	if ok = column.IsPrimaryKey(); !ok {
		t.Error("err")
	}
}

func Test_IsPrimaryKey_无(t *testing.T) {
	field, ok := getTestPersonFieldStruct("Name")
	if !ok {
		t.Error("err")
	}
	column := column{
		structOfField: field,
	}
	if ok = column.IsPrimaryKey(); ok {
		t.Error("err")
	}
}

func Test_IsForeignKey_有(t *testing.T) {
	field, ok := getTestPersonFieldStruct("AccountID")
	if !ok {
		t.Error("err")
	}
	column := column{
		structOfField: field,
	}
	if ok = column.IsForeignKey(); !ok {
		t.Error("err")
	}
}

func Test_IsForeignKey_无(t *testing.T) {
	field, ok := getTestPersonFieldStruct("Name")
	if !ok {
		t.Error("err")
	}
	column := column{
		structOfField: field,
	}
	if ok = column.IsForeignKey(); ok {
		t.Error("err")
	}
}

func getTestPersonFieldStruct(fieldName string) (reflect.StructField, bool) {
	entity := testPerson{}
	typeOfEntity := reflect.TypeOf(entity)
	return typeOfEntity.FieldByName(fieldName)
}
