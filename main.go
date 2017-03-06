package main

import (
	"os"

	"github.com/olsio/ftp-scan/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "My Cli Example"
	app.Usage = "Demo of urfave Cli Sweetness!"

	app.Commands = commands.Commands

	app.Run(os.Args)
}
