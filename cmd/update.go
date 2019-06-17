package cmd

import (
	l "log"

	"github.com/stqp/go-caniuse/pkg/datasrc"
	"github.com/urfave/cli"
)

func Update(c *cli.Context) (err error) {

	if _, err := datasrc.Update(true); err != nil {
		l.Fatal("Failed to update data source file")
		return err
	}

	l.Print("Successfulyy updated data source file")

	return nil
}
