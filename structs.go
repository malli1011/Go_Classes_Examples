package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string    `json:"name"`
	DoB       time.Time `json:"dob"`
	PhNumbers []string  `json:"phone omitempty"`
}

type tag1 struct {
	Name string
}

type tag2 struct {
	Name string
}

func structsDemo() {
	t1 := tag1{"Malli"}
	t2 := tag2{"Malli"}

	// t1==t2 // can't compare

	t3 := tag1(t2)

	fmt.Println(t3)
	fmt.Println(t1 == t3)

	m1 := struct {
		Name string `json:foo`
	}{"Malli"}

	m2 := struct {
		Name string `json:bar`
	}{"Arjun"}

	fmt.Println(m1, m2)

}
