package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "gotodo",
		Usage: "todo daemon and reminder",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "start",
				Usage:    "start the reminder daemon with the todo file",
				Aliases:  []string{"s"},
				FilePath: "$XDG_CONFIG_HOME/todo.json",
				Value:    "$XDG_CONFIG_HOME/todo.json",
				//EnvVars:  []string{"XDG_CONFIG_HOME"},
			},
		},
		Action: func(c *cli.Context) error {
			if filepath.Ext(c.Args().First()) == ".json" {

			} else {
				return cli.Exit("gotodo: configuration file must be json", 0)
			}

			if len(c.String("")) == 0 {

				fmt.Printf("starting in foreground with %s", c.Args().Get(0))
			} else if len(c.String("start")) > 0 {
				fmt.Printf("started in foreground with %s", c.String("start"))
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("")
	}
}
