package main

import (
	"os"
	"fmt"
	"math/rand"
	"strconv"
	"github.com/urfave/cli/v2"
	"github.com/fatih/color"
)

const PORT_MIN = 1024
const PORT_MAX = 65535

func main() {
	app := &cli.App{
		Name:  color.MagentaString("Random Port Generator")		,
		Usage: "Generate random and safe port numbers to use for local development",
		Commands: []*cli.Command{
			{
				Name:   "generate",
				Usage:  "generate a random port number",
				Action: func(cCtx *cli.Context) error {
					randomPort := generateRandomPort()
					prettyPortNumber := color.MagentaString(strconv.Itoa(randomPort))
					fmt.Printf("Generated random port:\n%s 🔌\n", prettyPortNumber)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func isPortReserved(port int) bool {
	return reservedPorts[port]
}

func generateRandomPort() int {
	for {
		port := rand.Intn(PORT_MAX - PORT_MIN + 1) + PORT_MIN
		if !isPortReserved(port) {
			return port
		}
	}
}