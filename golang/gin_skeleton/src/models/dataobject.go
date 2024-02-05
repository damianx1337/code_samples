package dataobject

type Tabler interface {
	TableName()	string
}

type TestTable struct {
	Field1	string	`gorm:"field1" json:"field1"`
}

// function used to redefine the table name as GORM will use test_table (camelcase -> underscore)
func (TestTable) TableName() string {
	return "dbschema.testtable"
}
