# go-dbfty

orm框架

## 配置

### 模型配置

```go
// 指定 ID值 | column 字段名、pk 主键 (任何值)、fk 外键 (任何值)
type Person struct{
    ID string `column:"id" pk:""`
    Name string `column:"name"`
    Age int `column:"age"`
    AccountID string `column:"account_id" fk:""`
}

// 自增ID
type PersonTwo struct{
    ID int `column:"id" auto:""`
    Name string
}

// TableName 固定结构体自定义表名函数
func (p Person) TableName() string {
    return "person"
}

```

## 示例

### 初始化

```go
// 读写分离使用
dbfty.SetDefault(
    mysql.Proxy(
        "读连接字段串",
        "写连接字段串",
    ),
)
dbfty.SetDefault(
    mysql.New("数据库连接字符串")
)
```

### 使用

```go
db := dbfty.Default().Db()
// entity 为数据的载体
entity := Person{
    ID: "test-person-001",
    Name: "Person001",
    Age: 1,
    AccountID:"test-account-id-001",
}
// 添加
// 1. 默认全字段添加
if err := db.Add(entity).Exec(); err != nil {
    return err
}
// 2. 指定添加字段
if err := db.Add(entity).Fields("Name", "Age").Exec(); err != nil {
    return err
}
// 删除
// 1. 默认主键删除
if err := db.Delete(entity).Exec(); err != nil {
    return err
}
// 2. 根据条件删除
if err := db.Delete(entity).Where("Name = ?", "Person002").Exec(); err != nil {
    return err
}
// 修改
// 1. 默认根据主键 => 全字段更新
if err := db.Update(entity).Exec(); err != nil {
    return err
}
// 2. 自定义条件更新、更新自定义字段
if err := db.Update(entity).Set("Name").Where("Name = ?", "Person001").Exec(); err != nil {
    return err
}
// 查询
// 查询单个数据
var result Person
if err := db.Query().Where("ID = ?", entity.ID).First(&result); err != nil {
    return err
}
fmt.Println("单个数据 => ", result)

// 查询所有数据
var results []Person
if err := db.Query().ToArray(&results); err != nil {
    return err
}
fmt.Println("查询person所有数据 => ", results)

// 还有相应的查询条件如 Order 升序, OrderByDesc 降序, GroupBy 分组, Take、Skip 用于分页
```

### 常规操作

