package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string] Vertex {
	"Zoho Corporation" : Vertex {
		40.82999, -73.97895,
	},
	"Google" : Vertex {
		37.42202, -122.08408,
	},
}

func main(){
	fmt.Println(m)
}