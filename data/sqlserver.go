package data

import (
	"extract-cli/config"
	"fmt"
)

type SQLServer struct {
	Database *config.Database
}

func NewSQLServer(d *config.Database) *SQLServer {
	return &SQLServer{
		Database: d,
	}
}

func (this *SQLServer) GetType() string {
	return "sqlserver"
}

func (this *SQLServer) GetString() string {
	// sqlserver://sa:mypass@localhost:1234?database=master&connection+timeout=30
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&connection+timeout=%d",
		this.Database.User, this.Database.Pass, this.Database.Host, this.Database.Port, this.Database.Schema, this.Database.Timeout)
}
