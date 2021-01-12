package dbfty

// IDelete 删除接口
type IDelete interface {
	Where(string, ...interface{}) IDelete
	Exec() error
}
