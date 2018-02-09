package main

import (
	"./git"
)

func drug() {
	git.Pull()
	write()
	git.Add()
	git.Commit()
	git.Push()
}
