package commands

import (
	"github.com/olsio/ftp-scan/store"
	"github.com/urfave/cli"
)

func List() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "list IP list from scan",
		Action: list,
	}
}

func list(c *cli.Context) error {
	store := store.Open()
	store.List()
	return store.Close()
}
