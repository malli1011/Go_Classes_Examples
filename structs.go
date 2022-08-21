package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name      string    `json:"name"`
	DoB       time.Time `json:"dob"`
	PhNumbers []string  `json:"phone,omitempty"`
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

	s2 := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	fmt.Println(m1, m2)
	p := Person{
		"Malli",
		s2,
		[]string{"9876543210"},
	}

	fmt.Println(p)

	b2, _ := json.Marshal(p)

	fmt.Println(string(b2))

	p2 := Person{
		"Malli",
		s2,
		[]string{},
	}

	fmt.Println(p2)

	b3, _ := json.Marshal(p2)
	fmt.Println(string(b3))

}
