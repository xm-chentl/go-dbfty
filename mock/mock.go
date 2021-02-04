package dbftymock

var (
	whereFunc func(entry interface{}, args ...interface{}) bool
)

// RegisterQuery 注册查询条件
func RegisterQuery(queryFunc func(entry interface{}, args ...interface{}) bool) {
	whereFunc = queryFunc
}

// Reset 重置
func Reset() {
	whereFunc = nil
}
