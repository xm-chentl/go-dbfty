package query

import (
	"fmt"
	"strings"

	dbfty "github.com/xm-chentl/go-dbfty"
	querycode "github.com/xm-chentl/go-dbfty/grammar/query/query-code"
)

// GroupBy 分组
type GroupBy struct {
	base
}

// Exec 执行
func (g GroupBy) Exec(ctx dbfty.IContext) {
	if ok := ctx.Has(querycode.GroupBy); ok {
		ctx.Set(
			querycode.SQL,
			fmt.Sprintf(
				" %s GROUP BY %s",
				ctx.Get(querycode.SQL).(string),
				strings.Join(
					ctx.Get(querycode.GroupBy).([]string),
					", ",
				),
			),
		)
	}

	g.base.Exec(ctx)
}
