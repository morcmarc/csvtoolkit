package main

import (
	"log"
	"os"

	. "github.com/morcmarc/csv2json/converter"

	"github.com/codegangsta/cli"
)

var convertCommand = cli.Command{
	Name:   "convert",
	Usage:  "Convert",
	Action: convertAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "",
			Usage: "path to file where result should be written to",
		},
	},
}

func convertAction(c *cli.Context) {
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

	// Create output file
	oFile := c.String("output")
	if oFile == "" {
		log.Fatalln("No output file")
	}
	json, err := os.Create(oFile)
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer json.Close()

	// New converter
	converter := NewConverter(csv, json)

	// Kick off the conversion
	converter.Run()
}
