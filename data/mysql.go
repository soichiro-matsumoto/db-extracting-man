package data

import (
	"extract-cli/config"
	"strconv"
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
	// user:password@tcp(host:3306)/schema"
	return this.Database.User +
		":" + this.Database.Pass +
		"@tcp(" + this.Database.Host +
		":" + strconv.Itoa(this.Database.Port) +
		")/" + this.Database.Schema
}
