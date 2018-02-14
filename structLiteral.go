package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X : 2} //Y = 0
	v3 = Vertex{} //X = 0, Y = 0
	p = &Vertex{3, 4} //has type *Vertex
)
func main() {
	fmt.Println(v1, v2, v3, p)
}