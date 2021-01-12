package grammar

// IBase 基础语法接口
type IBase interface {
	Generate(interface{}) (string, []interface{}, error)
}
