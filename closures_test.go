package main

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {

	f := Fib()
	g := Fib()

	for i := 1; i <= 10; i++ {
		fmt.Println(f(), g())
	}

}

func TestFuncWithBug(t *testing.T) {
	funcWithBug()
	funcWithBugFix()
}
