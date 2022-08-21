package main

import "fmt"

func Fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

func funcWithBug() {
	ss := make([]func(), 5)
	for i := 0; i < 5; i++ {
		ss[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}

	for i := 0; i < 5; i++ {
		ss[i]()
	}
}

func funcWithBugFix() {
	ss := make([]func(), 5)
	for i := 0; i < 5; i++ {
		i2 := i
		ss[i] = func() {
			fmt.Printf("%d @ %p\n", i2, &i2)
		}
	}

	for i := 0; i < 5; i++ {
		ss[i]()
	}
}
