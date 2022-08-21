package main

import (
	"fmt"
	"io"
	"os"
)

func Cat() {

	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		} else {
			_, err2 := io.Copy(os.Stdout, file)
			if err2 != nil {
				fmt.Fprintln(os.Stderr, err2)
				continue
			}
		}
		file.Close()
	}
}
