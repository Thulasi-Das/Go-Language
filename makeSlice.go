package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	s := a[0 : 3]
	printSlice(s)

	c := b[ : 5]
	printSlice(c)

	d := c[2 : 5]
	printSlice(d)
}
func printSlice(s []int){
	fmt.Printf("Len = %d, Cap = %d, Slice = %v\n", len(s), cap(s), s)
}