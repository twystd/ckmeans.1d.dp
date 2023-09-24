package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/twystd/ckmeans.1d.dp/go/dawq"
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
		fmt.Printf("\n  ckmeans-dawq %s\n\n", VERSION)
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
		fmt.Printf("  ... %v records read from %s\n", len(data), file)
	}

	clusters := dawq.CKMeans1dDp(data, nil, 1033, 1033)

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
		os.WriteFile(options.outfile, b.Bytes(), 0644)
	}
}

func read(f string) ([]dawq.Record, error) {
	data := []dawq.Record{}

	b, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bytes.NewReader(b))
	r.Comma = '\t'
	r.FieldsPerRecord = 2
	r.TrimLeadingSpace = true

	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		if len(row) == 2 {
			oid := row[0]
			if at, err := strconv.ParseFloat(row[1], 64); err != nil {
				fmt.Printf("   ... discarding row %v\n", row)
			} else {
				data = append(data, dawq.Record{
					OID: oid,
					At:  at,
				})
			}
		} else {
			fmt.Printf("   ... discarding row %v %v\n", row, len(row))
		}
	}

	return data, nil
}

func print(f io.Writer, clusters []dawq.Cluster) {
	for i, c := range clusters {
		line := fmt.Sprintf("%-4v", i+1)
		line += fmt.Sprintf(" %8.3f", c.Center)
		line += fmt.Sprintf(" %8.3f", c.Variance)
		for _, v := range c.Values {
			line += fmt.Sprintf(" [%-12v %-.3f]", v.OID, v.At)
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
