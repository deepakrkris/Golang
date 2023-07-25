package main

import (
	"fmt"
	"strings"
)

func countOccurs(s string, n int) int {
	length := len(s)
    countOfA := strings.Count(s, "a")
  
	totalInN := 0
	slots := n/length
	totalInN = slots * countOfA
 
	remainder := n%length

	return totalInN + strings.Count(s[:remainder], "a")
}

func main() {
    fmt.Println(countOccurs("abcac", 11))
	fmt.Println(countOccurs("abcac", 24))
	fmt.Println(countOccurs("abcaca", 24))
}
