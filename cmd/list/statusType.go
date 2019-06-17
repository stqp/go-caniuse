package list

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func Status(c *cli.Context) (err error) {

	header := []string{
		"support status", "description",
	}

	tableData := [][]string{
		{"Y", "(Y)es, supported by default."},
		{"A", "(A)lmost supported (Partially support)."},
		{"N", "(N)o support, or disabled by default."},
		{"P", "No support, but has (P)olyfill"},
		{"U", "Support (u)nknown."},
		{"X", "Requires prefi(x) to work."},
		{"D", "(D)isabled by default (need to enable flag or something)."},
		//{"#n", `Where n is a number, starting with 1, corresponds to the notes_by_num note. For example: "42":"y #1" means version 42 is supported by default and see note 1.`},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetCenterSeparator("*")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader(header)
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)
	table.AppendBulk(tableData)

	fmt.Println("")
	table.Render()
	fmt.Println("")

	return nil
}
