package dbfty

// IAdd 添加接口
type IAdd interface {
	// todo: v0.1.3版本 新增
	// Fields(...string) IAdd
	Exec() error
}
