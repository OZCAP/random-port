package commands

import (
	"fmt"
	"math/rand"
	"github.com/urfave/cli/v2"
	"github.com/fatih/color"
	"strconv"

	"github.com/OZCAP/random-port/config"
)

const PORT_MIN = 1024
const PORT_MAX = 65535

func GenerateAction() cli.ActionFunc {
	return func(cCtx *cli.Context) error {
		randomPort := generateRandomPort()
		prettyPortNumber := color.MagentaString(strconv.Itoa(randomPort))
		fmt.Printf("Generated random port:\n%s 🔌\n", prettyPortNumber)
		return nil
	}
}

func isPortReserved(port int) bool {
	return config.ReservedPorts[port]
}

func generateRandomPort() int {
	for {
		port := rand.Intn(PORT_MAX - PORT_MIN + 1) + PORT_MIN
		if !isPortReserved(port) {
			return port
		}
	}
}