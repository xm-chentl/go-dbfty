package dbfty

// IWriter 写接口
type IWriter interface {
	Add(interface{}) IAdd
	Delete(interface{}) IDelete
	Update(interface{}) IUpdate
}
