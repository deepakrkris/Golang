package main

import (
	"fmt"
)

type Vertex struct {
	lat, long float64
} 

func main() {
	var locations = map[string] Vertex {
		"Bell labs" : Vertex {
			90.99, 120.22,
		},
		"Tuticorin" : Vertex {
			22.00, 62.8002,
		},
	}

	fmt.Println(locations)
}
