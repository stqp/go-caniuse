package pkg

import (
	"fmt"
	l "log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/stqp/go-caniuse/pkg/datasrc"
	"github.com/stqp/go-caniuse/pkg/report"
	"github.com/stqp/go-caniuse/pkg/utils"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

func Run(c *cli.Context) (err error) {

	args := c.Args()
	if len(args) == 0 {
		l.Printf("go-canise require at least 1 argument.")
		cli.ShowAppHelpAndExit(c, 1)
	}

	search_key := args[0]

	json := datasrc.Data
	agents := gjson.GetBytes(json, "agents")
	data := gjson.GetBytes(json, "data")
	supported := data.Get(search_key)

	if !supported.Exists() {
		l.Printf("The feature seems not to be supported by any browser...")
		cli.ShowAppHelpAndExit(c, 1)
	}

	stats := supported.Get("stats")

	tableData := [][]string{
		{"NAME", "ID", "Y", "A", "N", "P", "X", "D", "U"},
	}

	stats.ForEach(func(browserId, versions gjson.Result) bool {

		statMap := map[byte][]string{}
		versions.ForEach(func(version, value gjson.Result) bool {
			k := value.String()[0]
			statMap[k] = append(statMap[k], version.String())
			return true
		})

		ks := []byte{'y', 'a', 'n', 'p', 'd', 'x', 'u'}

		for _, k := range ks {

			optimized := []string{}

			if len(statMap[k]) >= 3 {

				can_continue := true
				for _, version := range statMap[k] {
					if _, err := strconv.ParseFloat(version, 32); err != nil {
						can_continue = false
						break
					}
				}
				if can_continue == false {
					break
				}

				l, _ := strconv.ParseFloat(statMap[k][0], 32)
				u, _ := strconv.ParseFloat(statMap[k][1], 32)

				// do initially.
				if (u - l) > 1 {
					optimized = append(optimized, fmt.Sprint(l))
					l = u
				}

				// do main.
				for _, v := range statMap[k][2:] {
					version, _ := strconv.ParseFloat(v, 32)
					if (version - u) > 1 {
						if l == u {
							optimized = append(optimized, fmt.Sprint(u))
						} else {
							optimized = append(optimized, fmt.Sprint(l)+"~"+fmt.Sprint(u))
						}
						l = version
					}
					u = version
				}

				// do finally.
				if l == u {
					optimized = append(optimized, fmt.Sprint(u))
				} else {
					optimized = append(optimized, fmt.Sprint(l)+"~"+fmt.Sprint(u))
				}
			}

			statMap[k] = optimized
		}

		/*
			// I want to sort like this, but values are too messy.
			for _, k := range ks {
				sort.Slice(statMap[k], func(i, j int) bool {
					return statMap[k][i] < statMap[k][j]
				})
			}*/

		max_len := 0
		for _, k := range ks {
			if max_len < len(statMap[k]) {
				max_len = len(statMap[k])
			}
		}

		if c.String("browser") == "all" || c.String("browser") == browserId.String() {

			for i := 0; i < max_len; i++ {
				row := []string{}
				browser := agents.Get(browserId.String() + ".browser")
				row = append(row, browser.String())
				row = append(row, browserId.String())
				for _, k := range ks {
					if i < len(statMap[k]) {
						row = append(row, statMap[k][i])
					} else {
						row = append(row, "")
					}
				}
				tableData = append(tableData, row)
			}
		}

		return true
	})

	output := os.Stdout
	if c.String("output") != "" {
		workingdir, err := os.Getwd()
		if err != nil {
			l.Print(err)
			return nil
		}

		outpath := filepath.Join(workingdir, c.String("output"))
		exists, err := utils.Exists(outpath)
		if exists {
			l.Printf("Output file already exists: %s", outpath)
			return nil
		}

		outfile, err := os.Create(outpath)
		if err != nil {
			l.Printf("Failed to create output file: %s", outpath)
			l.Print(err)
			return nil
		}

		output = outfile
	}

	writer := report.NewWriter(c.String("format"), output)
	if err = writer.Write(tableData); err != nil {
		l.Fatalf("failed to write results: %w", err)
	}
	fmt.Println("")
	fmt.Println("INFO :")
	fmt.Println("  If you don't know much about status(= Y,A,N,P,X,D,U), try below command.")
	fmt.Println("    $ go-caniuse list status")
	fmt.Println("")

	return nil
}
