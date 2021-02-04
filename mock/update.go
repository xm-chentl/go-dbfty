package dbftymock

import (
	"fmt"
	"reflect"

	dbfty "github.com/xm-chentl/go-dbfty"
	"github.com/xm-chentl/go-dbfty/metadata"
)

type update struct {
	entry      interface{}
	repository *repository
	args       []interface{}
	sets       []string
}

func (u *update) Set(fields ...string) dbfty.IUpdate {
	u.sets = fields
	return u
}

func (u *update) Where(sql string, args ...interface{}) dbfty.IUpdate {
	u.args = args
	return u
}

func (u *update) Exec() error {
	table, err := metadata.Get(u.entry)
	if err != nil {
		return err
	}

	entries, ok := u.repository.data[table.Name()]
	if !ok {
		return nil
	}
	// 检查字段是否存在
	for _, field := range u.sets {
		if _, ok := table.GetColumnsByMap()[field]; !ok {
			return fmt.Errorf("table %s field (%s) is not exist", table.Name(), field)
		}
	}
	// 过滤并设置最新值
	for _, entry := range entries {
		if whereFunc != nil {
			if whereFunc(entry, u.args...) {
				for _, setField := range u.sets {
					column := table.GetColumnsByMap()[setField]
					fv, ok := column.GetValue(u.entry)
					if ok {
						reflect.ValueOf(
							entry,
						).Elem().FieldByName(
							column.GetStruct().Name,
						).Set(reflect.ValueOf(fv))
					}
				}
			}
		} else {
			// 只更新与u.entry主键有关的数据
			pk := table.GetPrimaryKeyBy()
			isOK := reflect.DeepEqual(
				reflect.ValueOf(entry).Elem().FieldByName(pk.GetStruct().Name).Interface(),
				reflect.ValueOf(u.entry).Elem().FieldByName(pk.GetStruct().Name).Interface(),
			)
			if isOK {
				for _, setField := range u.sets {
					column := table.GetColumnsByMap()[setField]
					fv, ok := column.GetValue(u.entry)
					if ok {
						reflect.ValueOf(
							entry,
						).Elem().FieldByName(
							column.GetStruct().Name,
						).Set(reflect.ValueOf(fv))
					}
				}
			}
		}
	}

	return nil
}
