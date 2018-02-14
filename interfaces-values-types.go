package main

import (
	"fmt"
	"math"
)

type I interface {
	print()
}

type T struct {
	S string
}

func (t *T) print() {
	fmt.Println(t.S)
}

type F float64

func (f F) print() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.print()

	i = F(math.Pi)
	describe(i)
	i.print()
}

func describe(i I){
	fmt.Printf("(%v, %T)\n", i, i)
}