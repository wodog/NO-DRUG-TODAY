package main

import (
  "os"
  "os/exec"
  "log"
  "time"
)


func main() {
  write()
  add()
  commit()
  push()
}

func write(){
  file, err := os.OpenFile("README.md", os.O_APPEND|os.O_WRONLY, 0644)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  byteSlice := []byte("\n\n" + getDate())
  w, err := file.Write(byteSlice)
  if err != nil {
    log.Fatal(err)
  }
  log.Println(w)
}

func getDate() string {
  return time.Now().String()
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
  if err !=nil {
    log.Fatal(err)
  }
}
