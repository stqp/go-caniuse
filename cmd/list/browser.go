package list

import (
	"fmt"
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
		{"no", "browser", "id"},
	}
	index := 1
	agents.ForEach(func(browserId, versions gjson.Result) bool {
		tableData = append(tableData, []string{
			strconv.Itoa(index), versions.Get("browser").String(), browserId.String(),
		})
		index++
		return true
	})

	output := os.Stdout
	writer := report.NewWriter(c.String("format"), output)
	fmt.Println("")
	if err = writer.Write(tableData); err != nil {
		l.Fatal("failed to write results: %w", err)
	}
	fmt.Println("")

	return nil
}
