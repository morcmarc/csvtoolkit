package main

import (
	"log"
	"os"

	. "github.com/morcmarc/csv2json/query"

	"github.com/codegangsta/cli"
)

var queryCommand = cli.Command{
	Name:   "query",
	Usage:  "Query CSV",
	Action: queryAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "query, q",
			Value: "",
			Usage: "query string",
		},
	},
}

func queryAction(c *cli.Context) {
	// Validate first argument which should be a file path
	if len(c.Args()) == 0 {
		log.Fatalln("No input given")
	}
	iFile := c.Args()[0]
	// Read file
	csv, err := os.Open(iFile)
	if err != nil {
		log.Fatalf("Could not open file: %s", iFile)
	}
	defer csv.Close()

	qs := c.String("query")
	if qs == "" {
		log.Fatalf("Missing query string\n")
	}

	q := NewQuery(csv)
	q.Run(qs)
}
