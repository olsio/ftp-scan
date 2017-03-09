package commands

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/olsio/ftp-scan/store"
	"github.com/olsio/ftp-scan/types"
	"github.com/urfave/cli"
)

func Scan() cli.Command {
	return cli.Command{
		Name:   "scan",
		Usage:  "scans IP",
		Action: scan,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "file",
				Value: "example.json",
			},
		},
	}
}

func scan(c *cli.Context) error {
	dbFile := c.String("file")
	file, err := os.Open(dbFile)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	store := store.Open()
	store.Clear()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var jsontype types.Grab
		json.Unmarshal([]byte(line), &jsontype)

		alreadyScanned, _ := store.Contains(jsontype.IP)

		if alreadyScanned {
			continue
		}

		if jsontype.Error == nil {
			if err := store.AddResult(jsontype.IP, types.Result{Directories: make([]string, 0)}); err != nil {
				log.Fatal(err)
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	return store.Close()
}
