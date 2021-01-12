package metadata

type testPerson struct {
	ID        string `pk:"true"`
	Name      string `column:"name"`
	Age       int    `column:"age"`
	AccountID string `column:"account_id" fk:"true"`
}

func (t testPerson) TableName() string {
	return "test-person"
}
