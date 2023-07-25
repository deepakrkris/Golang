package main

import (
	"fmt"
	"strings"
)

func find_unique_substring(str string) (int, string) {

	char_map := map[rune]int {}
	left := 0
	right := 0
	max_len := 0
	max_str := "" 

	for index, c := range str {
		if val, ok := char_map[c]; ok && val >= left {
			left = char_map[c] + 1
		}

		char_map[c] = index
		if (right - left + 1) > max_len {
			max_len = right - left + 1
			max_str = str[left : right + 1]
		}
		right += 1
	}

	return max_len, max_str
}

func test_data() []string {
	return []string {
		"abc",
		"abcabcbb",
		"bacbdefaccbb",
		"bacbdefacqer",
	}
}

func main() {
	for _, data := range test_data() {
		fmt.Println(strings.Repeat("_", 20))
		fmt.Println()
		fmt.Println("Result is ")
		fmt.Println(find_unique_substring(data))
		fmt.Println()
	}
}
