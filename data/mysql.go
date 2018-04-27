package data

import (
	"extract-cli/config"
	"fmt"
)

type MySQL struct {
	Database *config.Database
}

func NewMySQL(d *config.Database) *MySQL {
	return &MySQL{
		Database: d,
	}
}

func (this *MySQL) GetType() string {
	return "mysql"
}

func (this *MySQL) GetString() string {
	// user:password@tcp(host:port)/schema?timeout=30"
	return fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?timeout=%ds&charset=%s`,
		this.Database.User, this.Database.Pass, this.Database.Host, this.Database.Port, this.Database.Schema, this.Database.Timeout, this.Database.Charset)
}
