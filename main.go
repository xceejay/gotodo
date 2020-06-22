package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"path/filepath"
	//"syscall"
	//"unsafe"
)

//var TERMWIDTH uint = getWidth()

//var stars = make([]byte, TERMWIDTH)

//type winsize struct {
//Row    uint16
//Col    uint16
//Xpixel uint16
//Ypixel uint16
//}

//func getWidth() uint {
//ws := &winsize{}
//retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
//uintptr(syscall.Stdin),
//uintptr(syscall.TIOCGWINSZ),
//uintptr(unsafe.Pointer(ws)))

//if int(retCode) == -1 {
//panic(errno)
//}
//return uint(ws.Col)
//}

//func populateStars() []byte {

//}

func main() {
	notify()
	app := &cli.App{
		Name:  "gotodo",
		Usage: "todo daemon and reminder",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "start",
				Usage:   "start the reminder daemon with the todo file",
				Aliases: []string{"s"},
				Value:   "$XDG_CONFIG_HOME/todo.json",
				//EnvVars:  []string{"XDG_CONFIG_HOME"},
				TakesFile: true,
			},
		},
		Action: func(c *cli.Context) error {

			configfile := c.String("start")
			if filepath.Ext(c.Args().First()) == ".json" {
				if len(c.String("")) == 0 {
					fmt.Printf("starting in foreground with %s\n", c.Args().Get(0))
					start(c.Args().First())
				} else if len(c.String("start")) > 0 {
					fmt.Printf("started in foreground with %s\n", c.Args().First())
					start(configfile)
				}
			} else if filepath.Ext(c.String("start")) == ".json" {
				fmt.Printf("started in background with %s\n", c.String("start"))
				startservice(configfile)
			} else {
				return cli.Exit("gotodo: configuration file must be json\n", 0)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
