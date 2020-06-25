package main

import (
	"github.com/sevlyar/go-daemon"
	"log"
	"time"
)

func startservice(filename string) {

	dcontext := &daemon.Context{
		PidFileName: "gotodo.pid",
		PidFilePerm: 0644,
		LogFileName: "gotodo.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[gotodo]"},
	}

	d, err := dcontext.Reborn()
	if err != nil {
		log.Fatal(err)
	}
	if d != nil {
		return

	}
	defer dcontext.Release()

	log.Println("started in background")
	todolists := parse(filename)
	for {
		for i := 0; i < len(todolists.Todos); i++ {

			for timing, task := range todolists.Todos[i].Time {

				if time.Now().Format(datelayoutISO) == todolists.Todos[i].Date && timing == time.Now().Format(timelayoutISO) {

					notify(task)
				}

			}

		}
	}

}
