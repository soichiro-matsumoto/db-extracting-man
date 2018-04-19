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

	con, err := data.ConnectionCreate(db)
	if err != nil{
		return err
	}
	// query 実行
	client := data.NewDbClient(con)
	rows, err := client.Execute(args.Query)
	if err != nil {
		return err
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		return err
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
			return err
		}

		// 1行分
		r := []string{}
		// カラム数分ループ
		for i, col := range rawBytes {
			var value string
			if col != nil {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
			r = append(r, value)
		}
		// csv１行追加
		recordes = append(recordes, r)
	}

	helpers.ToCsv(args.OutputPath, recordes)
	return nil
}
