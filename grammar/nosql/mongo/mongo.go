package mongo

import "github.com/xm-chentl/go-dbfty/grammar"

type mongoImpl struct{}

func (m mongoImpl) Insert() grammar.IInsert{
	return nil
}

func (m mongoImpl) Delete() grammar.IDelete {
	return nil
}

func (m mongoImpl) Update() grammar.IUpdate {
	return nil
}

func (m mongoImpl) Select() grammar.ISelect{
	return nil
}

func New() grammar.IGrammar{
	return &mongoImpl{}
}
