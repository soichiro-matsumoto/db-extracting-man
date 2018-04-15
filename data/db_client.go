package data

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

type DbClient struct {
	Connection Connection
}

func NewDbClient(c Connection) *DbClient {
	return &DbClient{
		Connection: c,
	}
}

func (this DbClient) Execute(query string) (*sql.Rows, error) {
	// 接続
	// "sqlserver"の代わりに"mssql"でもOK
	con, err := sql.Open(this.Connection.GetType(), this.Connection.GetString())
	if err != nil {
		panic(err.Error())
	}

	// 切断
	defer con.Close()

	// select(複数行)
	rows, err := con.Query(query)

	return rows, err
}
