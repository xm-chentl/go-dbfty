package mysql

import "github.com/xm-chentl/go-dbfty/grammar"

type grammarMySQL struct{}

func (g grammarMySQL) Insert() grammar.IInsert {
	return &insert{}
}

func (g grammarMySQL) Delete() grammar.IDelete {
	return &delete{
		query: newWhere(),
	}
}

func (g grammarMySQL) Update() grammar.IUpdate {
	return &update{
		query: newWhere(),
	}
}

func (g grammarMySQL) Select() grammar.ISelect {
	return &selectSQL{
		query: newWhere(),
	}
}

// New 实例mysql语法生成器
func New() grammar.IGrammar {
	return &grammarMySQL{}
}
