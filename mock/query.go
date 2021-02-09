package dbftymock

import (
	"fmt"
	"reflect"

	dbfty "github.com/xm-chentl/go-dbfty"
	"github.com/xm-chentl/go-dbfty/metadata"
	"github.com/xm-chentl/go-dbfty/utils"
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
		defer func() {
			whereFunc = nil
		}()

		count = 0
		for _, entry := range entries {
			if whereFunc(entry) {
				count++
			}
		}
	}

	return count, nil
}

func (q query) First(entity interface{}) error {
	rt := reflect.TypeOf(entity)
	if rt.Kind() != reflect.Ptr {
		return fmt.Errorf("r is not ptr")
	}

	table, err := metadata.Get(entity)
	if err != nil {
		return err
	}

	entities, ok := q.repository.data[table.Name()]
	if !ok {
		return nil
	}

	if whereFunc != nil {
		defer Reset()

		for _, entity := range entities {
			if whereFunc(entity) {
				reflect.ValueOf(entity).Elem().Set(reflect.ValueOf(entity))
				break
			}
		}
	}

	return nil
}

func (q query) ToArray(entries interface{}) error {
	rt := reflect.TypeOf(entries)
	if rt.Kind() != reflect.Ptr {
		return fmt.Errorf("entities is not ptr")
	}

	typeOfEntity := utils.GetTypeBySlice(entries)
	table, err := metadata.Get(reflect.New(typeOfEntity).Interface())
	if err != nil {
		return err
	}

	entities, ok := q.repository.data[table.Name()]
	if !ok {
		return nil
	}

	if whereFunc != nil {
		defer Reset()

		results := reflect.MakeSlice(rt, 0, 0)
		for _, entity := range entities {
			if whereFunc(entity) {
				results = reflect.Append(results, reflect.ValueOf(entity))
				break
			}
		}
		reflect.ValueOf(entries).Elem().Set(results)
	}

	return nil
}

func (q query) Exc(entities interface{}, sql string, args ...interface{}) error {
	rt := reflect.TypeOf(entities)
	if rt.Kind() != reflect.Ptr {
		return fmt.Errorf("entities is not ptr")
	}

	typeOfEntity := utils.GetTypeBySlice(entities)
	table, err := metadata.Get(reflect.New(typeOfEntity).Interface())
	if err != nil {
		return err
	}

	entitiesOfTable, ok := q.repository.data[table.Name()]
	if !ok {
		return nil
	}

	if whereExcFunc != nil {
		defer Reset()

		results := reflect.MakeSlice(rt, 0, 0)
		for _, entity := range entitiesOfTable {
			if whereExcFunc(sql, args...) {
				results = reflect.Append(results, reflect.ValueOf(entity))
				break
			}
		}
		reflect.ValueOf(entities).Elem().Set(results)
	}

	return nil
}
