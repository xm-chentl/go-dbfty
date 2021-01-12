package query

import (
	"fmt"
	"strings"

	dbfty "github.com/xm-chentl/go-dbfty"
	querycode "github.com/xm-chentl/go-dbfty/grammar/query/query-code"
)

// OrderBy 排序
type OrderBy struct {
	base
}

// Exec 执行
func (o OrderBy) Exec(ctx dbfty.IContext) {
	orders := make([]string, 0)
	if ok := ctx.Has(querycode.OrderByAsc); ok {
		orders = append(
			orders,
			fmt.Sprintf(
				"%s ASC",
				strings.Join(
					ctx.Get(querycode.OrderByAsc).([]string),
					",",
				),
			),
		)
	}
	if ok := ctx.Has(querycode.OrderByDesc); ok {
		orders = append(
			orders,
			fmt.Sprintf(
				"%s DESC",
				strings.Join(
					ctx.Get(querycode.OrderByDesc).([]string),
					",",
				),
			),
		)
	}
	if len(orders) > 0 {
		ctx.Set(
			querycode.SQL,
			fmt.Sprintf(
				" %s ORDER BY %s",
				ctx.Get(querycode.SQL).(string),
				strings.Join(orders, ","),
			),
		)
	}

	o.base.Exec(ctx)
}
