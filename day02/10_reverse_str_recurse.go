package main

import (
	"fmt"
)

func reverse(str string) string {
	if len(str) == 0 {
		return ""
	}

	return reverse(str[1 : ]) + string(str[0])
}

func main() {
	fmt.Println(reverse("Hello Hi"))
}
