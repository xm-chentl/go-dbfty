package dbftymock

import (
	dbfty "github.com/xm-chentl/go-dbfty"
)

type repository struct {
	// key 表名 value 数据列表(实体)
	data map[string][]interface{}
}

func (r *repository) Ping() (bool, error) {
	panic("implement me")
}

func (r *repository) Add(entry interface{}) dbfty.IAdd {
	return &add{
		entry:      entry,
		repository: r,
	}
}

func (r *repository) Delete(entry interface{}) dbfty.IDelete {
	return nil
}

func (r *repository) Update(entry interface{}) dbfty.IUpdate {
	return nil
}

func (r *repository) Query() dbfty.IQuery {
	return &query{
		repository: r,
	}
}

func (r *repository) Ping() (bool, error) {
	return true, nil
}

func newRepository() *repository {
	return &repository{
		data: make(map[string][]interface{}),
	}
}
