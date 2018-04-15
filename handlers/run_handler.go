package handlers

import (
	"database/sql"
	"extract-cli/config"
	"extract-cli/data"
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

func RunHandler(c *cli.Context) error {

	// get args
	key := c.Args().Get(0)
	query := c.Args().Get(1)
	fmt.Println("key: " + c.Args().Get(0))
	fmt.Println("query: " + query)

	// config.tomlからDB接続情報を取得
	i, err := strconv.Atoi(key)
	if err != nil {
		return err
	}
	db := config.GetConfig().Databases[i]
	fmt.Println(
		"[" + strconv.Itoa(i) + "]\n" +
			db.ToString())

	// query 実行
	client := data.NewDbClient(data.NewMySQL(&db))
	rows, err := client.Execute(query)
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

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return err
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
	}

	// helpers.ToCsv()
	return nil
}
