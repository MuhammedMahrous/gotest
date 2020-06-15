package main

import (
	"golang.org/x/tour/wc"
	"strings"
//	"fmt"
)

func WordCount(s string) map[string]int {
	var wc = make(map[string] int)
	words := strings.Fields(s)
	for _,word := range words {
		if _,ok := wc[word]; ok {
			wc[word] = wc[word]+1
		} else {
			wc[word] = 1
		}
	}
//	fmt.Println(s,wc)
	return wc
}

func main() {
	wc.Test(WordCount)
}
