package dbftymock

var (
	whereFunc    func(entry interface{}, args ...interface{}) bool
	whereExcFunc func(sql string, args ...interface{}) bool
)

// RegisterQuery 注册查询条件
func RegisterQuery(queryFunc func(entry interface{}, args ...interface{}) bool) {
	whereFunc = queryFunc
}

// RegisterQueryExc 自定义查询语句
func RegisterQueryExc(queryExcFunc func(sql string, args ...interface{}) bool) {
	whereExcFunc = queryExcFunc
}

// Reset 重置
func Reset() {
	whereFunc = nil
	whereExcFunc = nil
}
