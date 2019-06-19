package list

import (
	l "log"
	"os"
	"strconv"

	"github.com/stqp/go-caniuse/pkg/datasrc"
	"github.com/stqp/go-caniuse/pkg/report"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

func Feature(c *cli.Context) (err error) {

	json := datasrc.Data

	data := gjson.GetBytes(json, "data")

	tableData := [][]string{
		{"NO", "FEATURE", "DESCRIPTION"},
	}
	index := 1
	data.ForEach(func(feature, detail gjson.Result) bool {
		tableData = append(tableData, []string{
			strconv.Itoa(index), feature.String(), detail.Get("description").String(),
		})
		index++
		return true
	})

	writer := report.NewWriter(c.String("format"), os.Stdout)
	if err = writer.Write(tableData); err != nil {
		l.Fatal("failed to write results: %w", err)
	}

	return nil
}
