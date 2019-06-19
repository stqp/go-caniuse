package report

import (
	"encoding/json"
	"fmt"
	"io"
	l "log"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Result [][]string

type Writer interface {
	Write(Result) error
}

type TableWriter struct {
	Output io.Writer
}

func (tw TableWriter) Write(result Result) error {
	table := tablewriter.NewWriter(tw.Output)
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetCenterSeparator("*")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)
	table.AppendBulk(result)
	fmt.Println("")
	table.Render()
	fmt.Println("")
	return nil
}

type CsvWriter struct {
	Output io.Writer
}

func (cw CsvWriter) Write(result Result) error {
	for _, row := range result {
		fmt.Fprintln(cw.Output, strings.Join(row, ","))
	}
	return nil
}

type JsonWriter struct {
	Output io.Writer
}

func (jw JsonWriter) Write(result Result) error {
	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		l.Fatal("failed to marshal json: %w", err)
		return err
	}

	if _, err = fmt.Fprint(jw.Output, string(output)); err != nil {
		l.Fatal("failed to write json: %w", err)
		return err
	}
	return nil
}

func NewWriter(format string, output *os.File) Writer {
	var writer Writer
	switch format {
	case "table":
		writer = TableWriter{Output: output}
	case "csv":
		writer = CsvWriter{Output: output}
	case "json":
		writer = JsonWriter{Output: output}
	}
	return writer
}
