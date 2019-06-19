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

func Browser(c *cli.Context) (err error) {

	json := datasrc.Data

	agents := gjson.GetBytes(json, "agents")

	tableData := [][]string{
		{"NO", "BROWSER", "ID"},
	}
	index := 1
	agents.ForEach(func(browserId, versions gjson.Result) bool {
		tableData = append(tableData, []string{
			strconv.Itoa(index), versions.Get("browser").String(), browserId.String(),
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
