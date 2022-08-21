package main

import "fmt"

var (
	a, b  int     = 10, 30
	c, d  float64 = 1.123, 3.123345
	s             = []int{1, 2, 3}
	chars         = []rune{'a', 'b', 'c'}
	m             = map[string]int{"one": 1, "two": 2}
)

func PrintFmt() {
	fmt.Printf("|%6d|%6d|\n", a, b)
	fmt.Printf("|%6.2f|%6.2f|\n", c, d)

	fmt.Println()

	fmt.Printf("slice with %%v: %v \n slice with %%#v %#v\n", s, s)

	fmt.Printf("%v\n", chars)
	fmt.Printf("%q\n", chars)

	fmt.Println()

	fmt.Printf("%T\n", m)
	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)
}
