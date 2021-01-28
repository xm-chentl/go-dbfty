package dbfty

// IFactory 数据工厂接口
type IFactory interface {
	Db() IRepository
	Uow() IUnitOfWork
	// todo: 预留接口用于组件做健康检测
	// IsHealth() (bool, error)
}
