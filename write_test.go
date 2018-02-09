package main

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	fmt.Println("aa")
}

func TestGetMessage(t *testing.T) {
	fmt.Println(getMessage())
}
