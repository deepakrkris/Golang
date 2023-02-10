package main

import (
	"fmt"
	"strings"
)

func slice_update(s []string) {
	if len(s) == 1 {
		s[0] = strings.ToUpper(s[0])
		return
	}
	slice_update(s[ : len(s) / 2])
	slice_update(s[len(s) / 2 :])
}

func main() {
	s := []string { "mala" , "vimala" , "vanaja", "kanaga" }
	slice_update(s)
	fmt.Println(s)

	t := []string { "akka" }

	for index, item := range s {
		s[index] = item + " " + t[0] + ","
	}

	fmt.Println(s)
}
