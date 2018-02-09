package main

import (
	"github.com/wodog/NO-DRUG-TODAY/git"
	"fmt"
)

func drug() {
	git.Pull()
	fmt.Println(1)
	write()
	fmt.Println(2)
	git.Add()
	fmt.Println(3)
	git.Commit()
	fmt.Println(4)
	git.Push()
	fmt.Println(5)
}
