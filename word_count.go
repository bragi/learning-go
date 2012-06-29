package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) (count map[string]int) {
	count = make(map[string]int)
	for _, v := range strings.Fields(s) {
		count[v] += 1
	}
	return 
}

func main() {
	wc.Test(WordCount)
}
