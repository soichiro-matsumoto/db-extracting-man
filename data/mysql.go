package data

import (
	"extract-cli/config"
	"fmt"
)

type MySQL struct {
	Type     string
	Database *config.Database
}

func NewMySQL(d *config.Database) *MySQL {
	return &MySQL{
		Type:     "mysql",
		Database: d,
	}
}

func (this *MySQL) GetType() string {
	return this.Type
}

func (this *MySQL) GetString() string {
	// user:password@tcp(host:port)/schema"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		this.Database.User, this.Database.Pass, this.Database.Host, this.Database.Port, this.Database.Schema)
}
