package dbfty

// Context 上下文内容
type Context map[int]interface{}

// Get 获取
func (c Context) Get(key int) interface{} {
	return c[key]
}

// Has 是否存在
func (c Context) Has(keys ...int) bool {
	isOk := true
	for _, key := range keys {
		if _, ok := c[key]; !ok {
			isOk = ok
			break
		}
	}

	return isOk
}

// Set 设置
func (c Context) Set(key int, value interface{}) {
	c[key] = value
}
