package main

import (
	"extract-cli/config"
	"extract-cli/handlers"
	"fmt"
	"os"
	"strconv"
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
		/*
			db
			configに記載されているDBのホスト名の一覧を表示する
		*/
		{
			Name:  "config",
			Usage: "config.tomlに設定されているDBの一覧を表示する",
			Action: func(c *cli.Context) error {
				fmt.Println("### config.tomlに設定されているDBの一覧を表示する")
				fmt.Println("/--------------------------------------/")
				for i, db := range config.GetConfig().Databases {
					fmt.Println(
						"[" + strconv.Itoa(i) + "]\n" +
							db.ToString())
				}

				fmt.Println("/--------------------------------------/")
				return nil
			},
		},
		{
			Name:   "csv",
			Usage:  "csv形式として出力する。csv [key] [sql_query] [output_path]",
			Action: handlers.RunHandler,
		},
		{
			Name:   "xml",
			Usage:  "xml形式として出力する。TODO:",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:   "json",
			Usage:  "JSON形式として出力する。TODO:",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:   "excel",
			Usage:  "Excel形式として出力する。TODO:",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	app.Run(os.Args)
}
