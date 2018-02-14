package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string] Vertex

func main() {
	m = make(map[string]Vertex)
	m["Zoho Corporation"] = Vertex{
		40.82999, -73.97895,
	}
	fmt.Println(m["Zoho Corporation"])
}