package dbftymock

import (
	"fmt"
	"testing"

	"github.com/xm-chentl/go-dbfty/metadata"
)

type testDeleteWork struct {
	ID      string `column:"id" pk:""`
	Content string
}

func Test_delete_Exec_Err(t *testing.T) {
	fmt.Println("del...1")
	testRepository := &repository{
		data: make(map[string][]interface{}),
	}
	table, err := metadata.Get(&testDeleteWork{})
	if err != nil {
		t.Error(err)
	}
	testRepository.data[table.Name()] = []interface{}{
		&testDeleteWork{
			ID:      "test001",
			Content: "test1",
		},
	}
	deleteOfRepository := &delete{
		entry:      &testDeleteWork{},
		repository: testRepository,
	}
	if err := deleteOfRepository.Where("content = ?", "test1").Exec(); err == nil {
		t.Error(err)
	}
}

func Test_delete_Exec(t *testing.T) {
	fmt.Println("del...2")
	testRepository := &repository{
		data: make(map[string][]interface{}),
	}
	table, err := metadata.Get(&testDeleteWork{})
	if err != nil {
		t.Error(err)
	}
	testRepository.data[table.Name()] = []interface{}{
		&testDeleteWork{
			ID:      "test001",
			Content: "test1",
		},
	}
	deleteOfRepository := &delete{
		entry:      &testDeleteWork{},
		repository: testRepository,
	}
	RegisterQuery(func(data interface{}, args ...interface{}) bool {
		entry := data.(*testDeleteWork)
		return entry.Content == args[0].(string)
	})
	if err := deleteOfRepository.Where("content = ?", "test1").Exec(); err != nil {
		t.Error(err)
	}
	if len(testRepository.data[table.Name()]) > 0 {
		t.Error("err")
	}
	// 重置
	Reset()
}
