package query

import (
	"fmt"

	dbfty "github.com/xm-chentl/go-dbfty"
	querycode "github.com/xm-chentl/go-dbfty/grammar/query/query-code"
)

// Limit 分页
type Limit struct {
	base
}

// Exec 执行
func (l Limit) Exec(ctx dbfty.IContext) {
	if ok := ctx.Has(querycode.Take); ok {
		skip := 0
		if ok = ctx.Has(querycode.Skip); ok {
			skip = ctx.Get(querycode.Skip).(int)
		}
		take := ctx.Get(querycode.Take).(int)
		ctx.Set(
			querycode.SQL,
			fmt.Sprintf(
				" %s LIMIT %d, %d",
				ctx.Get(querycode.SQL).(string),
				skip, take,
			),
		)
	}
}
