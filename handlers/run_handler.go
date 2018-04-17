package handlers

import (
	"database/sql"
	"extract-cli/config"
	"extract-cli/data"
	"extract-cli/helpers"
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

type Args struct {
	Key        string
	Query      string
	OutputPath string
}

func NewArgs(c *cli.Context) *Args {
	path := c.Args().Get(2)
	if path == "" {
		path = "./output.csv"
	}
	return &Args{
		Key:        c.Args().Get(0),
		Query:      c.Args().Get(1),
		OutputPath: path,
	}
}

func (this *Args) GetDatabase(c *config.Config) (*config.Database, error) {
	i, err := strconv.Atoi(this.Key)
	if err != nil {
		return nil, err
	}
	db := c.Databases[i]
	return &db, nil
}

func RunHandler(c *cli.Context) error {

	args := NewArgs(c)
	fmt.Println("Args")
	fmt.Println(args)

	// config.tomlからDB接続情報を取得
	db, err := args.GetDatabase(config.GetConfig())
	if err != nil {
		return err
	}
	fmt.Println("Database")
	fmt.Println(db.ToString())

	// query 実行
	client := data.NewDbClient(data.NewMySQL(db))
	rows, err := client.Execute(args.Query)
	if err != nil {
		return err
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		return err
	}
	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	recordes := [][]string{}
	for rows.Next() {
		r := []string{}
		err = rows.Scan(scanArgs...)
		if err != nil {
			return err
		}

		for i, col := range values {
			var value string
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
			// csv１行追加
			r = append(r, value)
		}

		recordes = append(recordes, r)
	}

	helpers.ToCsv(args.OutputPath, recordes)
	return nil
}
