package data

import (
	"extract-cli/config"
	"errors"
	"fmt"
)

type Connection interface {
	GetType() string

	GetString() string
}

func NewConnection(db *config.Database) (Connection, error){
	switch db.System {
	case "mysql":
		return NewMySQL(db), nil
	case "sqlserver":
		return NewSQLServer(db), nil
	default:
		return nil, errors.New(fmt.Sprintf("not defined type : %s", db.System))
	}
}