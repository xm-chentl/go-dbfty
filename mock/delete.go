package dbftymock

import (
	"fmt"

	dbfty "github.com/xm-chentl/go-dbfty"
	"github.com/xm-chentl/go-dbfty/metadata"
)

type delete struct {
	entry      interface{}
	args       []interface{}
	repository *repository
}

func (d *delete) Where(sql string, args ...interface{}) dbfty.IDelete {
	d.args = args
	return d
}

func (d *delete) Exec() error {
	table, err := metadata.Get(d.entry)
	if err != nil {
		return err
	}

	entries, ok := d.repository.data[table.Name()]
	if !ok {
		return nil
	}

	if whereFunc == nil {
		return fmt.Errorf("mock where is not register => mock.RegisterQuery(...)")
	}

	for index, entry := range entries {
		if whereFunc(entry, d.args...) {
			d.repository.data[table.Name()] = append(entries[0:index], entries[index+1:]...)
		}
	}

	return nil
}
