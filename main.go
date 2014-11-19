package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "csvtoolkit"
	app.Usage = "convert, query csv files"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		convertCommand,
		queryCommand,
	}
	app.Run(os.Args)
}
