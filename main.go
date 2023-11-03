package main

import (
	"os"
	"github.com/urfave/cli/v2"
	"github.com/fatih/color"

	"github.com/OZCAP/random-port/commands"
)

func main() {
	app := &cli.App{
		Name:  color.MagentaString("Random port generator")		,
		Usage: "Generate random and safe port numbers to use for local development",
		Commands: []*cli.Command{
			{
				Name:   "generate",
				Usage:  "generate a random port number",
				Action: commands.GenerateAction(),
			},
		},
		Flags: []cli.Flag{
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
