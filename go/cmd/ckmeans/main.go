package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/twystd/ckmeans.1d.dp/go/ckmeans"
)

const VERSION = "v0.0.0"

var options = struct {
	outfile string
	debug   bool
}{
	outfile: "",
	debug:   false,
}

func main() {
	flag.StringVar(&options.outfile, "out", options.outfile, "output file path")
	flag.BoolVar(&options.debug, "debug", options.debug, "enables debugging")
	flag.Parse()

	if options.debug {
		fmt.Printf("\n  ckmeans.1d.dp %s\n\n", VERSION)
	}

	if len(flag.Args()) == 0 {
		usage()
		os.Exit(1)
	}

	file := flag.Args()[0]
	if options.debug {
		fmt.Printf("  ... reading data from %s\n", file)
	}

	data, err := read(file)
	if err != nil {
		fmt.Printf("\n  ** ERROR: unable to read data from file %s (%v)\n\n", file, err)
		os.Exit(1)
	}

	if options.debug {
		fmt.Printf("  ... %v values read from %s\n", len(data), file)
	}

	clusters := ckmeans.CKMeans1dDp(data, nil, func(v any) float64 { return v.(float64) })

	if options.debug {
		fmt.Printf("  ... %v clusters\n", len(clusters))
	}

	var b bytes.Buffer

	print(&b, clusters)

	if options.outfile == "" {
		fmt.Println()
		fmt.Printf("%s", string(b.Bytes()))
		fmt.Println()
	} else {
		ioutil.WriteFile(options.outfile, b.Bytes(), 0644)
	}
}

func read(f string) ([]any, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	data := []any{}
	tokens := regexp.MustCompile(`\s+`).Split(string(b), -1)

	for _, t := range tokens {
		if t != "" {
			if v, err := strconv.ParseFloat(t, 64); err != nil {
				fmt.Printf("  ** WARN: invalid value (%s)\n", t)
			} else {
				data = append(data, v)
			}
		}
	}

	return data, nil
}

func print(f io.Writer, clusters []ckmeans.Cluster) {
	columns := 0
	for _, c := range clusters {
		if len(c.Values) > columns {
			columns = len(c.Values)
		}
	}
	columns += 3

	table := make([][]string, len(clusters))
	for i := range table {
		table[i] = make([]string, columns)
	}

	for i, c := range clusters {
		table[i][0] = fmt.Sprintf("%d", i+1)
		table[i][1] = fmt.Sprintf("%v", c.Center)
		table[i][2] = fmt.Sprintf("%v", c.Variance)
		for j, v := range c.Values {
			table[i][j+3] = fmt.Sprintf("%v", v)
		}
	}

	widths := make([]int, columns)
	for _, row := range table {
		for i, s := range row {
			if len(s) > widths[i] {
				widths[i] = len(s)
			}
		}
	}

	formats := make([]string, columns)
	for i, w := range widths {
		formats[i] = fmt.Sprintf("%%-%dv", w)
	}

	for i, c := range clusters {
		line := ""
		line += fmt.Sprintf(formats[0], i+1)
		line += "  "
		line += fmt.Sprintf(formats[1], c.Center)
		line += "  "
		line += fmt.Sprintf(formats[2], c.Variance)
		line += "  "
		for j, v := range c.Values {
			line += " "
			line += fmt.Sprintf(formats[j+3], v)
		}

		fmt.Fprintf(f, "%s\n", line)
	}
}

func usage() {
	fmt.Println()
	fmt.Println("  Usage: ckmeans [options] <file>")
	fmt.Println()
	fmt.Println("  Arguments:")
	fmt.Println()
	fmt.Println("    file  Path to file containing the whitespace delimited data to be clustered")
	fmt.Println()
	fmt.Println("  Options:")
	fmt.Println()
	fmt.Println("    --debug     Displays internal information for diagnosing errors")
	fmt.Println()
}
