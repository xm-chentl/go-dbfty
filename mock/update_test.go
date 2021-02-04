package dbftymock

import (
	"testing"

	"github.com/xm-chentl/go-dbfty/metadata"
)

type testUpdateWork struct {
	ID      string `column:"id" pk:""`
	Content string
}

func Test_update_Exec(t *testing.T) {
	testRepository := &repository{
		data: make(map[string][]interface{}),
	}
	table, err := metadata.Get(testUpdateWork{})
	if err != nil {
		t.Error(err)
	}
	testRepository.data[table.Name()] = []interface{}{
		&testUpdateWork{
			ID:      "test001",
			Content: "test1",
		},
		&testUpdateWork{
			ID:      "test002",
			Content: "test2",
		},
	}
	updateOfRepository := &update{
		entry: &testUpdateWork{
			ID:      "test002",
			Content: "test03",
		},
		repository: testRepository,
	}
	if err := updateOfRepository.Set("content").Exec(); err != nil {
		t.Error(err)
	}
	for _, dataItem := range testRepository.data[table.Name()] {
		entry := dataItem.(*testUpdateWork)
		if entry.ID == "test002" && entry.Content != "test03" {
			t.Error("err")
		}
	}

	// 重置
	Reset()
}

func Test_update_Exec_where(t *testing.T) {
	testRepository := &repository{
		data: make(map[string][]interface{}),
	}
	table, err := metadata.Get(testUpdateWork{})
	if err != nil {
		t.Error(err)
	}
	testRepository.data[table.Name()] = []interface{}{
		&testUpdateWork{
			ID:      "test001",
			Content: "test1",
		},
		&testUpdateWork{
			ID:      "test002",
			Content: "test2",
		},
	}
	updateOfRepository := &update{
		entry: &testUpdateWork{
			Content: "test03",
		},
		repository: testRepository,
	}
	RegisterQuery(func(item interface{}, args ...interface{}) bool {
		entry := item.(*testUpdateWork)
		return entry.ID == args[0].(string)
	})
	if err := updateOfRepository.Set("content").Where("id = ?", "test001").Exec(); err != nil {
		t.Error(err)
	}
	for _, dataItem := range testRepository.data[table.Name()] {
		entry := dataItem.(*testUpdateWork)
		if entry.ID == "test001" && entry.Content != "test03" {
			t.Error("err")
		}
	}

	// 重置
	Reset()
}
