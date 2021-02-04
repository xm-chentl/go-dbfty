package dbftymock

import (
	"testing"
)

type tesAddtWork struct {
	ID      string `column:"id" pk:""`
	Content string
}

func Test_add_Exec(t *testing.T) {
	testRepository := &repository{
		data: make(map[string][]interface{}),
	}
	addOfRepository := &add{
		entry: &tesAddtWork{
			ID:      "test-001",
			Content: "test",
		},
		repository: testRepository,
	}
	if err := addOfRepository.Exec(); err != nil {
		t.Error("err", err)
	}
	if len(testRepository.data) > 1 {
		t.Error("err")
	}
	// 重置
	Reset()
}
