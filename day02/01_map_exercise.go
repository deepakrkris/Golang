package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
    words := strings.Split(s, " ")
 	count_map := map[string]int{}

	for _, w := range words {
	   if len(strings.TrimSpace(w)) > 0 {
    	  count_map[w] += 1
	   }
	}
	return count_map
}

func main() {
	wc.Test(WordCount)
}
