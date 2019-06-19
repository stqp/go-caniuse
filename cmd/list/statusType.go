package list

import (
	l "log"
	"os"

	"github.com/stqp/go-caniuse/pkg/report"
	"github.com/urfave/cli"
)

func Status(c *cli.Context) (err error) {

	tableData := [][]string{
		{"SUPPORT STATUS", "DESCRIPTION"},
		{"Y", "(Y)es, supported by default."},
		{"A", "(A)lmost supported (Partially support)."},
		{"N", "(N)o support, or disabled by default."},
		{"P", "No support, but has (P)olyfill"},
		{"U", "Support (u)nknown."},
		{"X", "Requires prefi(x) to work."},
		{"D", "(D)isabled by default (need to enable flag or something)."},
		//{"#n", `Where n is a number, starting with 1, corresponds to the notes_by_num note. For example: "42":"y #1" means version 42 is supported by default and see note 1.`},
	}

	writer := report.NewWriter(c.String("format"), os.Stdout)
	if err = writer.Write(tableData); err != nil {
		l.Fatal("failed to write results: %w", err)
	}

	return nil
}
