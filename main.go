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
	app.Name = "extract"
	app.Usage = "クエリを実行してCSVに出力します。"
	app.Version = "0.0.1"
	app.Compiled = time.Now()

	app.Commands = []cli.Command{
		/*
			db
			configに記載されているDBのホスト名の一覧を表示する
		*/
		{
			Name:  "db",
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
		/*
			run [key] [sql-path] [output-path]
		*/
		{
			Name:   "run",
			Usage:  "抽出実行する",
			Action: handlers.RunHandler,
		},
	}

	app.Run(os.Args)
}
