package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/metakeule/fmtdate"
)

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
