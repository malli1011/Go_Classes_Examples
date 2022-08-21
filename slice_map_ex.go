package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var words = make(map[string]int)

type kv struct {
	key string
	val int
}

func main() {

	scan := bufio.NewScanner(os.Stdin)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "Uniqure words")

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss[:5] {
		fmt.Println(s.key, "appears", s.val, "times")
	}
}
