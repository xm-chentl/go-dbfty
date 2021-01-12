package dbfty

// IReader 读接口
type IReader interface {
	// 支持单表查询
	Query() IQuery
	// 可扩展联表
	// ... Join、Left、Right
}
