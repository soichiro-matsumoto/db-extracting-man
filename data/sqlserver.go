package data

import (
	"extract-cli/config"
	"fmt"
)

type SQLServer struct {
	Type     string
	Database *config.Database
}

func NewSQLServer(d *config.Database) *SQLServer {
	return &SQLServer{
		Type:     "sqlserver",
		Database: d,
	}
}

func (this *SQLServer) GetType() string {
	return this.Type
}

func (this *SQLServer) GetString() string {
	// sqlserver://sa:mypass@localhost:1234?database=master&connection+timeout=30
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&connection+timeout=30",
		this.Database.User, this.Database.Pass, this.Database.Host, this.Database.Port, this.Database.Schema)
}
