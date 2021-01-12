package mysql

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/xm-chentl/go-dbfty/grammar"
	"github.com/xm-chentl/go-dbfty/grammar/sql"
	"github.com/xm-chentl/go-dbfty/metadata"
	"github.com/xm-chentl/go-dbfty/utils"
)

type insert struct {
	fields []string
}

func (i *insert) Fields(fields ...string) grammar.IInsert {
	i.fields = fields

	return i
}

func (i *insert) Generate(data interface{}) (string, []interface{}, error) {
	table, err := metadata.Get(data)
	if err != nil {
		return "", nil, err
	}

	columns := make([]metadata.IColumn, 0)
	if len(i.fields) == 0 {
		columns = table.GetColumns()
	} else {
		for _, field := range i.fields {
			if column, ok := table.GetColumnsByMap()[field]; ok {
				if column.IsAuto() {
					continue
				}
				columns = append(columns, column)
			}
		}
	}

	args := make([]interface{}, 0)
	valueOfData := reflect.ValueOf(data)
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		valueOfData = valueOfData.Elem()
	}
	for _, column := range columns {
		args = append(
			args,
			valueOfData.FieldByName(column.GetStruct().Name).Interface(),
		)
	}

	i.fields = make([]string, 0)
	for _, column := range columns {
		i.fields = append(i.fields, column.Field())
	}

	return fmt.Sprintf(
		sql.FORMATSQLINSERT,
		table.Name(),
		strings.Join(i.fields, ","),
		strings.Join(
			utils.GenSep(len(i.fields), "?"), ",",
		),
	), args, nil
}
