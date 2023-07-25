package main

import (
	"fmt"
)

func reverse(str string) string {
	strArray := []rune(str)
	i := 0
	j := len(str) - 1
	for i < j {
		strArray[i], strArray[j] = strArray[j], strArray[i]
		i += 1
		j -= 1
	}

	return string(strArray)
}

func main() {
	str := "Hello Hi"
	fmt.Println(reverse(str)) 
}
