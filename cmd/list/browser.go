package list

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/stqp/go-caniuse/pkg/datasrc"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

func Browser(c *cli.Context) (err error) {

	json := datasrc.Data

	agents := gjson.GetBytes(json, "agents")

	header := []string{
		"no", "browser", "id",
	}

	tableData := [][]string{}
	index := 1
	agents.ForEach(func(browserId, versions gjson.Result) bool {
		tableData = append(tableData, []string{
			strconv.Itoa(index), versions.Get("browser").String(), browserId.String(),
		})
		index++
		return true
	})

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
