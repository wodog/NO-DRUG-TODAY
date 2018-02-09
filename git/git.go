package git

import (
	"log"
	"os/exec"
)

func Add() {
	cmd := exec.Command("git", "add", ".")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Commit() {
	cmd := exec.Command("git", "commit", "-m", "add")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Push() {
	cmd := exec.Command("git", "push", "origin", "master:master")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Pull() {
	cmd := exec.Command("git", "pull", "origin", "master")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
