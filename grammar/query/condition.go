package query

import (
	"fmt"

	dbfty "github.com/xm-chentl/go-dbfty"
	querycode "github.com/xm-chentl/go-dbfty/grammar/query/query-code"
)

// Condition 条件
type Condition struct {
	base
}

// Exec 执行
func (w Condition) Exec(ctx dbfty.IContext) {
	if ok := ctx.Has(querycode.Condition); ok {
		ctx.Set(
			querycode.SQL,
			fmt.Sprintf(
				" WHERE %s",
				ctx.Get(querycode.Condition).(string),
			),
		)
		if args := ctx.Get(querycode.ConditionArgs).([]interface{}); args != nil && len(args) > 0 {
			ctx.Set(
				querycode.Args,
				ctx.Get(querycode.ConditionArgs).([]interface{}),
			)
		}
	} else {
		ctx.Set(querycode.SQL, "")
	}

	w.base.Exec(ctx)
}
