package grammar

// IGrammar  语法接口
type IGrammar interface {
	Insert() IInsert
	Delete() IDelete
	Update() IUpdate
	Select() ISelect
}
