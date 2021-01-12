package dbfty

import (
	"sync"
)

// DEFAULT 默认key
const DEFAULT = "default"

var (
	errHandler   func(err error)
	keyOfFactory = make(map[string]IFactory)
	mx           sync.Mutex
)

// Default 获取数据库连接实例
func Default() IFactory {
	mx.Lock()
	defer mx.Unlock()

	db, ok := keyOfFactory[DEFAULT]
	if !ok {
		panic("未设置默认仓储实例")
	}

	return db
}

// Get 获取指定key的仓储实例
func Get(key string) IFactory {
	mx.Lock()
	defer mx.Unlock()

	return keyOfFactory[key]
}

// Set 设置相应key的仓储实例
func Set(key string, db IFactory) {
	mx.Lock()
	defer mx.Unlock()

	keyOfFactory[key] = db
}

// SetDefault 设计默认key的仓储实例
func SetDefault(db IFactory) {
	mx.Lock()
	defer mx.Unlock()

	keyOfFactory[DEFAULT] = db
}

// ErrHandler 错误处理
func ErrHandler(f func(error)) {
	errHandler = f
}
