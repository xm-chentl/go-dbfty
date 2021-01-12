package query

import (
	dbfty "github.com/xm-chentl/go-dbfty"
	querycode "github.com/xm-chentl/go-dbfty/grammar/query/query-code"
)

type base struct {
	nextHandler IHandler
}

func (b *base) Next(handler IHandler) IHandler {
	b.nextHandler = handler
	return b.nextHandler
}

func (b base) Exec(ctx dbfty.IContext) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		ctx.Set(querycode.Err, err)
	// 	}
	// }()

	if ok := ctx.Has(querycode.Err); ok {
		return
	}
	if b.nextHandler != nil {
		b.nextHandler.Exec(ctx)
	}
}
