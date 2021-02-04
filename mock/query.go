package dbftymock

import (
	dbfty "github.com/xm-chentl/go-dbfty"
	"github.com/xm-chentl/go-dbfty/metadata"
)

type query struct {
	repository *repository
}

func (q *query) Order(fields ...string) dbfty.IQuery {
	return q
}

func (q *query) OrderByDesc(fields ...string) dbfty.IQuery {
	return q
}

func (q *query) GroupBy(fields ...string) dbfty.IQuery {
	return q
}

func (q *query) Take(num int) dbfty.IQuery {
	return q
}

func (q *query) Skip(num int) dbfty.IQuery {
	return q
}

func (q *query) Where(where string, args ...interface{}) dbfty.IQuery {
	return q
}

func (q query) Count(entry interface{}) (int, error) {
	table, err := metadata.Get(entry)
	if err != nil {
		return 0, err
	}

	entries, ok := q.repository.data[table.Name()]
	if !ok {
		return 0, nil
	}

	count := len(entries)
	if whereFunc != nil {
		count = 0
		for _, entry := range entries {
			if whereFunc(entry) {
				count++
			}
		}
	}

	return count, nil
}

func (q query) First(entry interface{}) error {
	return nil
}

func (q query) ToArray(entries interface{}) error {
	return nil
}
