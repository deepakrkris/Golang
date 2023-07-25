package main

import (
    "fmt"
	"strings"
)

func rev_words(str string) string {
	words := strings.Split(str, " ")
	result := []string{}

	for _, w := range words {
		if len(strings.TrimSpace(w)) > 0 {
			result = append([]string { strings.TrimSpace(w) }, result...)
		}
	}

	return strings.Join(result[:], " ")
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
