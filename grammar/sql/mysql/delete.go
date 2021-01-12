package mysql

import (
	"fmt"

	"github.com/xm-chentl/go-dbfty/grammar"
	"github.com/xm-chentl/go-dbfty/grammar/sql"
	"github.com/xm-chentl/go-dbfty/metadata"
)

type delete struct {
	query grammar.IWhere
}

func (d *delete) Query() grammar.IWhere {
	return d.query
}

func (d *delete) Generate(data interface{}) (string, []interface{}, error) {
	table, err := metadata.Get(data)
	if err != nil {
		return "", nil, err
	}
	if !d.query.IsLegal() {
		// 条件缺失
		value, _ := table.GetPrimaryKeyBy().GetValue(data)
		d.query.Where(
			fmt.Sprintf("`%s` = ?", table.GetPrimaryKeyBy().Name()),
			value,
		)
	}

	// 自定义条件
	where, whereArgs, err := d.query.Generate(data)
	if err != nil {
		return "", nil, err
	}

	return fmt.Sprintf(
		sql.FORMATSQLDELETE,
		table.Name(),
		where,
	), whereArgs, nil
}
