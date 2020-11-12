package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/twystd/ckmeans.1d.dp/go/ckmeans"
)

const VERSION = "v0.0.0"

var options = struct {
	debug bool
}{
	debug: false,
}

func main() {
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

	clusters := ckmeans.CKMeans(data, nil)

	if options.debug {
		fmt.Printf("  ... %v clusters\n", len(clusters))
	}

	fmt.Println()
	print(clusters)
	fmt.Println()
}

func read(f string) ([]float64, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	data := []float64{}
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

func print(clusters [][]float64) {
	columns := 0
	for _, c := range clusters {
		if len(c) >= columns {
			columns = len(c) + 1
		}
	}

	table := make([][]string, len(clusters))
	for i := range table {
		table[i] = make([]string, columns)
	}

	for i, c := range clusters {
		table[i][0] = fmt.Sprintf("%d", i+1)
		for j, v := range c {
			table[i][j+1] = fmt.Sprintf("%v", v)
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
		for j, v := range c {
			line += " "
			line += fmt.Sprintf(formats[j+1], v)
		}
		fmt.Printf("%s\n", line)
	}
}

// func format(array []float64) string {
// 	var b bytes.Buffer
// 	for _, v := range array {
// 		fmt.Fprintf(&b, "%0.6f ", v)
// 	}
//
// 	return strings.TrimSpace(string(b.Bytes()))
// }

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
