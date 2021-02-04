package dbftymock

import "github.com/xm-chentl/go-dbfty/metadata"

type add struct {
	entry      interface{}
	repository *repository
}

func (a *add) Exec() error {
	table, err := metadata.Get(a.entry)
	if err != nil {
		return err
	}
	if _, ok := a.repository.data[table.Name()]; !ok {
		a.repository.data[table.Name()] = make([]interface{}, 0)
	}
	a.repository.data[table.Name()] = append(a.repository.data[table.Name()], a.entry)

	return nil
}
