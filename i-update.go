package dbfty

// IUpdate 更新接口
type IUpdate interface {
	Set(...string) IUpdate
	Where(string, ...interface{}) IUpdate
	Exec() error
}
