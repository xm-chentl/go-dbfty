package grammar

// IInsert 插入语法接口
type IInsert interface {
	IBase
	Fields(...string) IInsert
}
