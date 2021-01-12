package grammar

// IDelete 删除语法接口
type IDelete interface {
	IBase
	Query() IWhere
}
