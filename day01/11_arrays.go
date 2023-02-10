package main

import (
	"fmt"
)

func main() {
	var a [2]int;
	var b [2]string;

	a[0] = 10
	b[0] = "hello"
	b[1] = "world"

	fmt.Println(a, b)
}
