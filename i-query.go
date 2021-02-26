package dbfty

// IQuery 查询接口
type IQuery interface {
	// 可扩展为聚合函数
	// Aggregates(...IAggregateFunc) IQuery // 自定义聚合函数
	// Fields(...string) IQuery             // 自定义查询字段
	Order(...string) IQuery
	OrderByDesc(...string) IQuery
	GroupBy(...string) IQuery
	Take(int) IQuery // 获取
	Skip(int) IQuery // 跳过
	Where(string, ...interface{}) IQuery
	Count(interface{}) (int, error)
	First(interface{}) error
	ToArray(interface{}) error
	Exc(interface{}, string, ...interface{}) error
}
