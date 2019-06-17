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

func Feature(c *cli.Context) (err error) {

	json := datasrc.Data

	data := gjson.GetBytes(json, "data")

	header := []string{
		"feature", "Description",
	}

	tableData := [][]string{}
	index := 1
	data.ForEach(func(feature, detail gjson.Result) bool {
		tableData = append(tableData, []string{
			strconv.Itoa(index), feature.String(), detail.Get("description").String(),
		})
		fmt.Println(detail.Get("description").String())
		index++
		return false
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
