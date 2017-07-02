package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/metakeule/fmtdate"
	"github.com/robfig/cron"
)

func main() {
	ch := make(chan int)
	c := cron.New()
	c.AddFunc("0 22 17 * * *", drug)
	c.Start()
	<-ch
}

func drug() {
	pull()
	write()
	add()
	commit()
	push()
}

func pull() {
	cmd := exec.Command("git", "pull")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func write() {
	dir, _ := os.Getwd()
	filePath := filepath.Join(dir, "README.md")
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Write([]byte(getMessage()))
}

func getMessage() string {
	date := fmtdate.Format("YYYY-MM-DD", time.Now())
	return date + ": NO DRUG TODAY!\n\n"
}

func add() {
	cmd := exec.Command("git", "add", ".")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func commit() {
	cmd := exec.Command("git", "commit", "-m", "add")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func push() {
	cmd := exec.Command("git", "push", "origin", "master:master")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
