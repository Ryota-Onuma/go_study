package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m := map[string]string{
		"name": "ryota",
		"address": "yokohama",
	}
	fmt.Println(m["name"])
	map_practice := make(map[string]int, 3)
	map_practice["a"] = 1
	map_practice["b"] = 2
	map_practice["c"] = 3
	map_practice["d"] = 4
	fmt.Println(map_practice)
	_, ok := map_practice["d"]
	_, ok2 := map_practice["e"]
	fmt.Println(ok,ok2)
}
