package querycode

const (
	// Data 数据实现
	Data = iota
	// Fields 指定字段
	Fields
	// Condition 条件
	Condition
	// ConditionArgs 条件参数
	ConditionArgs
	// Take 取多少
	Take
	// Skip 跳过多少
	Skip
	// GroupBy 分组
	GroupBy
	// OrderByAsc 正序
	OrderByAsc
	// OrderByDesc 倒序
	OrderByDesc
	// SQL 语句
	SQL
	// Args 参数
	Args
	// Err 错误
	Err
)
