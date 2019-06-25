package main

import (
	"log"
	"os"

	"github.com/stqp/go-caniuse/cmd"
	"github.com/stqp/go-caniuse/cmd/list"
	"github.com/stqp/go-caniuse/pkg"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "go-caniuse"
	app.Usage = "show browser support status table for web technologies."
	app.Action = pkg.Run
	app.Version = "dev"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "format, f",
			Value: "table",
			Usage: "format (table, csv, json)",
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

	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "list something.",
			Subcommands: []cli.Command{
				{
					Name:   "browser",
					Usage:  "list all browsers.",
					Action: list.Browser,
					Flags:  flags,
				},
				{
					Name:   "feature",
					Usage:  "list all feature.",
					Action: list.Feature,
					Flags:  flags,
				},
				{
					Name:   "status",
					Usage:  "list support status flags.",
					Action: list.Status,
					Flags:  flags,
				},
			},
		},
		{
			Name:   "update",
			Usage:  "update data source file",
			Action: cmd.Update,
		},
	}

	app.Flags = flags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
