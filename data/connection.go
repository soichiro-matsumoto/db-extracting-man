package data

type Connection interface {

	// sqlserver
	GetType() string

	// sqlserver://ohke:p@ssw0rd@testdb.ikyu.jp?database=TestDB
	GetString() string
}
