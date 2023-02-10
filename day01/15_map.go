package main

import (
	"fmt"
)

type Vertex struct {
	lat, long float64
}


func main() {
	var map_to_string map[string] Vertex

	map_to_string = make(map[string] Vertex)

	map_to_string["Bell Labs"] = Vertex { lat : 40.80 , long : 66.90 }

	fmt.Println(map_to_string)
}
