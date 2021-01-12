package dbfactory

// IRepository 仓储接口
type IRepository interface {
	IWriter
	IReader
}
