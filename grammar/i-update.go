package grammar

// IUpdate 更新语法接口
type IUpdate interface {
	IBase
	Set(...string) IUpdate
	Query() IWhere
}
