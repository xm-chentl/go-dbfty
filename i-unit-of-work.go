package dbfty

// IUnitOfWork 工作单元接口
type IUnitOfWork interface {
	Commit() error
}
