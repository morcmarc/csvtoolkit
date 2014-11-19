package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "csv2json"
	app.Usage = "Convert csv files into json objects."
	app.Version = "0.1.1"
	app.Commands = []cli.Command{
		convertCommand,
		queryCommand,
	}
	app.Run(os.Args)
}
