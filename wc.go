package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WC() {
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int
		f, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(f)
		for scan.Scan() {
			lc++

			s := scan.Text()

			cc += len(s)

			wc += len(strings.Fields(s))

			fmt.Printf("%6d %6d %6d %s", lc, wc, cc, fname)
		}

		f.Close()
	}
}
