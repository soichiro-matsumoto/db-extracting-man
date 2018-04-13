package data

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func execute(c Connection, query string) (*sql.Rows, error) {
	// 接続
	// "sqlserver"の代わりに"mssql"でもOK
	con, err := sql.Open(c.GetType(), c.GetString())
	if err != nil {
		panic(err.Error())
	}

	// 切断
	defer con.Close()

	// select(複数行)
	rows, err := con.Query(query)

	return rows, err
}
