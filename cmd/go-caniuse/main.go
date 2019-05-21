package main

import(
	"log"
	"os"
	"github.com/urfave/cli"

	"github.com/stqp/trivy/pkg"
	"github.com/stqp/trivy/pkg/log"
)
var(
	version = "dev"
)

func main(){

	app :=cli.NewApp()
	app.Name = "go-caniuse"
	app.Version = version
	//app.ArgsUsage = "image_name"
	//app.Usage = "A simple and comprehensive vulnerability scanner for containers"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "format, f",
			Value: "table",
			Usage: "format (table, json)",
		}
	}

	app.Action = pkg.Run

	err := app.Run(os.Args)
	if err != nil {
		if log.Logger != nil {
			log.Fatal(err)
		}
		l.Fatal(err)
	}

}