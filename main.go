package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"fmt"

	"github.com/robfig/cron"
	"github.com/metakeule/fmtdate"
)

func main() {
	drug()
	ch := make(chan int)
	c := cron.New()
	c.AddFunc("0 22 17 * * *", drug)
	c.Start()
	<-ch
}

func drug() {
	fmt.Println("pull...")
	pull()
	fmt.Println("write...")
	write()
	fmt.Println("add...")
	add()
	fmt.Println("commit...")
	commit()
	fmt.Println("push...")
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
