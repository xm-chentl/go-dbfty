package dbfty

// IContext 上下文
type IContext interface {
	Get(int) interface{}
	Has(...int) bool
	Set(int, interface{})
}
