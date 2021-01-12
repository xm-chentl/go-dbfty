package mysql

import (
	"fmt"
	"strings"

	"github.com/xm-chentl/go-dbfty/grammar"
	"github.com/xm-chentl/go-dbfty/grammar/sql"
	"github.com/xm-chentl/go-dbfty/metadata"
)

const ()

type update struct {
	sets  []string
	query grammar.IWhere
}

func (u *update) Set(fields ...string) grammar.IUpdate {
	if len(fields) > 0 {
		u.sets = append(u.sets, fields...)
	}

	return u
}

func (u *update) Query() grammar.IWhere {
	return u.query
}

func (u update) Generate(data interface{}) (string, []interface{}, error) {
	table, err := metadata.Get(data)
	if err != nil {
		return "", nil, nil
	}

	args := make([]interface{}, 0)
	setFields := make([]string, 0)
	if len(u.sets) > 0 {
		for _, field := range u.sets {
			if column, ok := table.GetColumnsByMap()[field]; ok {
				setFields = append(setFields, fmt.Sprintf("%s = ?", column.Field()))
				value, _ := column.GetValue(data)
				args = append(args, value)
			} else {
				return "", nil, fmt.Errorf(`sql grammar (update) err: "%s" field no exist`, field)
			}
		}
	} else {
		for _, column := range table.GetColumns() {
			if value, isNil := column.GetValue(data); isNil {
				args = append(args, value)
				setFields = append(setFields, fmt.Sprintf("%s = ?", column.Field()))
			}
		}
	}

	if !u.query.IsLegal() {
		// 条件是否合法
		value, _ := table.GetPrimaryKeyBy().GetValue(data)
		u.query.Where(
			fmt.Sprintf("`%s` = ?", table.GetPrimaryKeyBy().Name()),
			value,
		)
	}

	where, whereArgs, err := u.query.Generate(data)
	if err != nil {
		return "", nil, err
	}
	args = append(args, whereArgs...)

	return fmt.Sprintf(
		sql.FORMATSQLUPDATE,
		table.Name(),
		strings.Join(setFields, ","),
		where,
	), args, nil
}
