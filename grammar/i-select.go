package grammar

// ISelect 查询语法接口
type ISelect interface {
	IBase
	Fields(...string) ISelect
	Aggregation(...IAggregation) ISelect
	Query() IWhere
}
