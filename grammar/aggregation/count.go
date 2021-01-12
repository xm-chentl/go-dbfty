package aggregation

import (
	"fmt"

	"github.com/xm-chentl/go-dbfty/grammar"
)

type count struct {
	field string
}

func (c count) Generate() string {
	return fmt.Sprintf("COUNT(%s)", c.field)
}

// Count COUNT聚合函数
func Count(field string) grammar.IAggregation {
	return &count{
		field: field,
	}
}
