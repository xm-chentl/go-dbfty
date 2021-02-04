package dbftymock

import dbfty "github.com/xm-chentl/go-dbfty"

type factory struct {
	repository *repository
}

func (f *factory) Db() dbfty.IRepository {
	if f.repository == nil {
		f.repository = newRepository()
	}
	return f.repository
}

func (f *factory) Uow() dbfty.IUnitOfWork {
	return nil
}

// New 新建一个mock实例
func New() dbfty.IFactory {
	return nil
}
