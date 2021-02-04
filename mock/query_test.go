package dbftymock

import (
	"fmt"
	"testing"

	"github.com/xm-chentl/go-dbfty/metadata"
)

type testQueryWork struct {
	ID      string `column:"id" pk:""`
	Content string
}

func Test_query_Count_All(t *testing.T) {
	fmt.Println("query...1")
	testRepository := &repository{
		data: make(map[string][]interface{}),
	}
	table, err := metadata.Get(testQueryWork{})
	if err != nil {
		t.Error(err)
	}
	testRepository.data[table.Name()] = []interface{}{
		&testQueryWork{
			ID:      "test-001",
			Content: "test1",
		},
		&testQueryWork{
			ID:      "test002",
			Content: "test2",
		},
	}
	testQuery := &query{
		repository: testRepository,
	}
	total, err := testQuery.Count(testQueryWork{})
	if err != nil {
		t.Error(err)
	}
	if total != 2 {
		t.Error("total is not equal 2")
	}
}

func Test_query_Count_Query(t *testing.T) {
	fmt.Println("query...2")
	testRepository := &repository{
		data: make(map[string][]interface{}),
	}
	table, err := metadata.Get(testQueryWork{})
	if err != nil {
		t.Error(err)
	}
	testRepository.data[table.Name()] = []interface{}{
		&testQueryWork{
			ID:      "test-001",
			Content: "test1",
		},
		&testQueryWork{
			ID:      "test002",
			Content: "test2",
		},
	}
	testQuery := &query{
		repository: testRepository,
	}
	RegisterQuery(func(data interface{}, _ ...interface{}) bool {
		entry := data.(*testQueryWork)
		return entry.ID == "test002"
	})
	total, err := testQuery.Count(testQueryWork{})
	if err != nil {
		t.Error(err)
	}
	if total != 1 {
		t.Error("total is not equal 1")
	}

	// 重置
	Reset()
}
