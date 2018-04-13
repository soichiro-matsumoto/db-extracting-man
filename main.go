package main

import (
	"extract/config"
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
				fmt.Println("/--------------------------------------/")
				fmt.Println("config.tomlに設定されているDBの一覧を表示する")

				for i, v := range config.GetConfig().Databases {
					fmt.Println(
						"[" + strconv.Itoa(i) + "]\n" +
							"	Host:" + v.Host + "\n" +
							"	System:" + v.System + "\n" +
							"	Encoding:" + v.Encoding)
				}

				fmt.Println("/--------------------------------------/")
				return nil
			},
		},
		/*
			run [key]
		*/
		{
			Name:  "run",
			Usage: "抽出実行する",
			Action: func(c *cli.Context) error {

				key := c.Args().Get(0)
				fmt.Println("key：" + arg1)

				arg2 := c.Args().Get(1)
				fmt.Println("2つ目：" + arg2)

				db := config.GetConfig().Databases[strconv.Atoi(key)]

				return nil
			},
		},
	}

	app.Run(os.Args)
}
