package mysql

import (
	"fmt"
	"strings"

	"github.com/xm-chentl/go-dbfty/grammar"
	"github.com/xm-chentl/go-dbfty/grammar/sql"
	"github.com/xm-chentl/go-dbfty/metadata"
)

type selectSQL struct {
	fields       []string
	aggregations []grammar.IAggregation
	query        grammar.IWhere
}

func (s *selectSQL) Fields(fields ...string) grammar.ISelect {
	s.fields = fields

	return s
}

func (s *selectSQL) Aggregation(aggregations ...grammar.IAggregation) grammar.ISelect {
	s.aggregations = aggregations

	return s
}

func (s *selectSQL) Query() grammar.IWhere {
	return s.query
}

func (s *selectSQL) Generate(data interface{}) (string, []interface{}, error) {
	table, err := metadata.Get(data)
	if err != nil {
		return "", nil, nil
	}

	if len(s.fields) > 0 {
		for _, field := range s.fields {
			if _, ok := table.GetColumnsByMap()[field]; !ok {
				return "", nil, fmt.Errorf("sql grammar (select) err: %s no exist", field)
			}
		}
	} else {
		s.fields = make([]string, 0)
		for _, field := range table.GetColumns() {
			s.fields = append(s.fields, field.Name())
		}
	}

	columns := s.toColumns()
	// todo... 聚合（与其它字段是互斥的），后期根据场景进行优化
	if len(s.aggregations) > 0 {
		columns = make([]string, 0)
		for _, aggregation := range s.aggregations {
			columns = append(columns, aggregation.Generate())
		}
	}

	// 允许查询所有 | 可能需要取消，正常使用场景一般不会需要，后续使用中看情况
	if !s.query.IsLegal() {
		return fmt.Sprintf(
			sql.FORMATSQLSELECT,
			strings.Join(columns, ","),
			table.Name(),
			"",
		), nil, nil
	}

	whereOfSQL, args, err := s.query.Generate(data)
	if err != nil {
		return "", nil, err
	}

	return fmt.Sprintf(
		sql.FORMATSQLSELECT,
		strings.Join(columns, ","),
		table.Name(),
		whereOfSQL,
	), args, nil
}

func (s *selectSQL) toColumns() []string {
	fields := make([]string, 0)
	for _, field := range s.fields {
		fields = append(fields, fmt.Sprintf("`%s`", field))
	}

	return fields
}
