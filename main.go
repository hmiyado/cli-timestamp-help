// code from https://cli.urfave.org/v2/#timestamp-flag
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.TimestampFlag{Name: "meeting", Layout: "2006-01-02T15:04:05"},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Printf("%s", cCtx.Timestamp("meeting").String())
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
