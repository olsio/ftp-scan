package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/olsio/ftp-scan/store"
	"github.com/olsio/ftp-scan/types"
	"github.com/urfave/cli"
)

func Load() cli.Command {
	return cli.Command{
		Name:   "load",
		Usage:  "loads IP list from scan",
		Action: load,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "file",
				Value: "example.json",
			},
		},
	}

}

func load(c *cli.Context) error {
	dbFile := c.String("file")
	file, err := os.Open(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	store := store.NewStore("my.db")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		var jsontype types.Grab
		json.Unmarshal([]byte(line), &jsontype)
		fmt.Printf("Results: %v\n", jsontype)

		if jsontype.Error == nil {
			store.AddScanTarget([]byte(line))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
