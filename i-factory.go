package dbfty

// IFactory 数据工厂接口
type IFactory interface {
	Db() IRepository
	Uow() IUnitOfWork
}
