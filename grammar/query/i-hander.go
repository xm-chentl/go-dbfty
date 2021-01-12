package query

import dbfty "github.com/xm-chentl/go-dbfty"

// IHandler 处理接口
type IHandler interface {
	Next(IHandler) IHandler
	Exec(dbfty.IContext)
}
