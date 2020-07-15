package main

import (
	"github.com/urfave/cli"
)

const usage = "A File Download App"

func main() {
	app := cli.NewApp()
	app.Name = "fdl"
	app.Usage = usage
	app.Commands = []cli.Command{
		clientCommand,
		serverCommand,
	}
}
