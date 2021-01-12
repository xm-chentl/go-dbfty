package mysql

import (
	"github.com/xm-chentl/go-dbfty/grammar"
	"github.com/xm-chentl/go-dbfty/grammar/query"
	querycode "github.com/xm-chentl/go-dbfty/grammar/query/query-code"
)

type where struct {
	ctx     dbfactory.IContext
	handler query.IHandler

	customSQL  string
	customArgs []interface{}
}

func (w *where) Fields(fields ...string) grammar.IWhere {
	return w
}

func (w *where) GroupBy(fields ...string) grammar.IWhere {
	w.ctx.Set(querycode.GroupBy, fields)

	return w
}

func (w *where) OrderBy(fields ...string) grammar.IWhere {
	w.ctx.Set(querycode.OrderByDesc, fields)

	return w
}

func (w *where) OrderByDesc(fields ...string) grammar.IWhere {
	w.ctx.Set(querycode.OrderByDesc, fields)

	return w
}

func (w *where) Take(num int) grammar.IWhere {
	w.ctx.Set(querycode.Take, num)

	return w
}

func (w *where) Skip(num int) grammar.IWhere {
	w.ctx.Set(querycode.Skip, num)

	return w
}

func (w *where) Where(sql string, args ...interface{}) grammar.IWhere {
	w.ctx.Set(querycode.Condition, sql)
	w.ctx.Set(querycode.ConditionArgs, args)

	return w
}

func (w where) IsLegal() bool {
	return w.ctx.Has(querycode.Condition) && w.ctx.Get(querycode.Condition).(string) != ""
}

func (w where) Generate(data interface{}) (string, []interface{}, error) {
	w.handler.Exec(w.ctx)
	ok := false
	if ok = w.ctx.Has(querycode.Err); ok {
		return "", nil, w.ctx.Get(querycode.Err).(error)
	}
	if ok := w.ctx.Has(querycode.SQL, querycode.Args); ok {
		return w.ctx.Get(querycode.SQL).(string), w.ctx.Get(querycode.Args).([]interface{}), nil
	}

	return w.ctx.Get(querycode.SQL).(string), nil, nil
}

func newWhere() grammar.IWhere {
	handler := new(query.Condition)
	handler.Next(
		new(query.GroupBy),
	).Next(
		new(query.OrderBy),
	).Next(
		new(query.Limit),
	)

	return &where{
		ctx:     dbfactory.Context{},
		handler: handler,
	}
}
