package main

import (
	"github.com/urfave/cli/v2"
)

var FlushDBCmd = &cli.Command{
	Name:        "flushdb",
	Description: "truncate relevant tables in the DB",
	Usage:       "clair-load-test flushdb",
	Action:      flushDBAction,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "override",
			Aliases: []string{"y"},
			Usage:   "do not ask to confirm",
			Value:   false,
			EnvVars: []string{"_OVERRIDE"},
		},
	},
}

func flushDBAction(c *cli.Context) error {
	return nil
}
