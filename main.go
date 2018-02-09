package main

import (
	"log"
	"time"

	"github.com/sevlyar/go-daemon"
)

func main() {
	context := &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
	}
	p, err := context.Reborn()
	if err != nil {
		log.Fatalln(err)
	}
	if p != nil {
		return
	}

	drug()
	ticker := time.NewTicker(24 * time.Hour)
	for _ = range ticker.C {
		drug()
	}
}
