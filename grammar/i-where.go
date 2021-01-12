package grammar

// IWhere 条件语法接口
type IWhere interface {
	IBase
	GroupBy(...string) IWhere
	OrderBy(...string) IWhere
	OrderByDesc(...string) IWhere
	Take(int) IWhere
	Skip(int) IWhere
	Where(string, ...interface{}) IWhere
	IsLegal() bool
}
