package handlers

import (
	"database/sql"
	"extract-cli/config"
	"extract-cli/data"
	"extract-cli/helpers"
	"fmt"

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

func CsvHandler(c *cli.Context) error {

	fmt.Println("start ....")

	args := NewArgs(c)

	// config.tomlからDB接続情報を取得
	db, err := config.GetConfig().GetDatabase(args.Key)
	if err != nil {
		return cli.NewExitError(err, 500)
	}
	fmt.Println("selected database")
	fmt.Println(db.ToString())

	// Connection 生成
	con, err := data.NewConnection(db)
	if err != nil{
		return cli.NewExitError(err, 500)
	}

	// query 実行
	client := data.NewDbClient(con)
	rows, err := client.Execute(args.Query)
	if err != nil {
		return cli.NewExitError(err, 500)
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		return cli.NewExitError(err, 500)
	}
	rawBytes := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.
	scanArgs := make([]interface{}, len(rawBytes))
	for i := range rawBytes {
		scanArgs[i] = &rawBytes[i]
	}

	// csv出力するの連想配列
	recordes := [][]string{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return cli.NewExitError(err, 500)
		}

		// 1行分
		r := []string{}
		// カラム数分ループ
		for _, col := range rawBytes {
			var value string
			if col != nil {
				value = string(col)
			}
			r = append(r, value)
		}
		// csv１行追加
		recordes = append(recordes, r)
	}

	err = helpers.ToCsv(args.OutputPath, recordes)
	if err != nil {
		return cli.NewExitError(err, 500)
	}

	fmt.Println("end ....")

	return nil
}
