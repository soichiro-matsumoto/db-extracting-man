package main

import (
	"extract-cli/config"
	"extract-cli/handlers"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "extract-cli"
	app.Usage = "クエリを実行してCSVに出力します。"
	app.Version = "0.0.1"
	app.Compiled = time.Now()

	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "config.tomlに設定されているDBの一覧を表示する",
			Action: func(c *cli.Context) error {
				fmt.Println("### config.tomlに設定されているDBの一覧を表示する")
				fmt.Println("/--------------------------------------/")
				for _, db := range config.GetConfig().Databases {
					fmt.Println(db.ToString())
				}

				fmt.Println("/--------------------------------------/")
				return nil
			},
		},
		{
			Name:   "csv",
			Usage:  "csv形式として出力する。csv [config_key] [sql_filepath] [output_path]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "query, q",
					Usage: "queryを直接指定",
				},
			},
			Action: handlers.CsvHandler,
		},
		{
			Name:   "xml",
			Usage:  "xml形式として出力する。TODO:",
			Action: func(c *cli.Context) error {
				fmt.Println("Unimplemented ....")
				return nil
			},
		},
		{
			Name:   "json",
			Usage:  "JSON形式として出力する。TODO:",
			Action: func(c *cli.Context) error {
				fmt.Println("Unimplemented ....")
				return nil
			},
		},
		{
			Name:   "excel",
			Usage:  "Excel形式として出力する。TODO:",
			Action: func(c *cli.Context) error {
				fmt.Println("Unimplemented ....")
				return nil
			},
		},
	}

	app.Run(os.Args)
}
