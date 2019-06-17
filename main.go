package main

import (
	"log"
	"os"

	"github.com/stqp/go-caniuse/cmd"
	"github.com/stqp/go-caniuse/cmd/list"
	"github.com/stqp/go-caniuse/pkg"
	"github.com/urfave/cli"
)

/*
・data.jsonの場所を指定する機能
    web or file
    $ go-caniuse websockets -f data.json
・ブラウザ一覧表示機能
    略称
    $ go-caniuse list browser
  機能一覧表示機能
  	$ go-caniuse list feature
  	$ go-caniuse list type
・絞り込み検索
    ブラウザ略称
    キーワード曖昧検索
    $ go-caniuse -b <browser_key> websockets
・出力指定
    csv, table
    os.Sysout, file
    $ go-caniuse -f [csv|table] websockets
  	$ go-caniuse -o <file_path> websockets
・ヘルプの充実
    yanpuxdとかの意味を
*/

func main() {

	app := cli.NewApp()
	app.Name = "go-caniuse"
	app.Usage = "show browser support status table for front-end technologies."
	app.Action = pkg.Run
	app.Version = "dev"

	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "list something.",
			Subcommands: []cli.Command{
				{
					Name:   "browser",
					Usage:  "list browsers.",
					Action: list.Browser,
				},
				{
					Name:   "feature",
					Usage:  "list browser's feature.",
					Action: list.Feature,
				},
				{
					Name:   "status",
					Usage:  "list browser's support status.",
					Action: list.Status,
				},
			},
		},
		{
			Name:   "update",
			Usage:  "update data source file",
			Action: cmd.Update,
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "format, f",
			Value: "table",
			Usage: "format (table, json)",
		},
		cli.StringFlag{
			Name:  "input, i",
			Value: "",
			Usage: "data file path",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "",
			Usage: "output file name",
		},
		cli.StringFlag{
			Name:  "browser, b",
			Value: "all",
			Usage: "browser name.",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
