package commands

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	Scan(),
	List(),
}
