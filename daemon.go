package main

import (
	"github.com/sevlyar/go-daemon"
	"log"
)

func startservice(file string) {

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

}
