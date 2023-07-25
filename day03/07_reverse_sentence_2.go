package main

import (
    "fmt"
	"strings"
)

func rev_words(str string) string {
	words := strings.Split(str, " ")
	result := ""

	for _, w := range words {
		if len(strings.TrimSpace(w)) > 0 {
			result = strings.TrimSpace(w) + "  " + result
		}
	}

	return result
}

func test_data () []string {
	return []string {
		"hello    world         to you",
		"world hello",
	}
}

func main () {
	for _, data := range test_data() {
		fmt.Println(rev_words(data))
	}
}
