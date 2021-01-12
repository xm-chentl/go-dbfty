package metadata

import (
	"reflect"
	"sync"
)

var (
	keyOfTable = make(map[string]ITable)
	mx         sync.Mutex
)

// Get 获取实体的反射数据 => 隐患并发问题 优化方案 sync.xxx读写锁
func Get(entity interface{}) (ITable, error) {
	mx.Lock()
	defer mx.Unlock()

	// 获取结构体
	typeOfEntity := reflect.TypeOf(entity)
	valueOfEntity := reflect.ValueOf(entity)
	if typeOfEntity.Kind() == reflect.Ptr {
		typeOfEntity = typeOfEntity.Elem()
		valueOfEntity = valueOfEntity.Elem()
	}

	key := typeOfEntity.Name()
	// 判断是否有缓存
	if table, ok := keyOfTable[key]; ok {
		return table, nil
	}

	// 缓存
	keyOfTable[key] = &table{
		typeOfData:  typeOfEntity,
		valueOfDate: valueOfEntity,
	}
	if err := keyOfTable[key].Check(); err != nil {
		return nil, err
	}

	return keyOfTable[key], nil
}
