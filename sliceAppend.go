package main

import "fmt"

func main() {
	// var s []int
	a := [5]int{11, 12, 13, 14, 15}
	s := a[ : ]
	printSlice(s)

	s =append(s, 0)
	printSlice(s)

	s = append(s, 1, 2, 3, 4)
	printSlice(s)

	s = s[2 : 4]
	printSlice(s)

	s = append(s, 5)
	printSlice(s)

	s = append(s, 6, 7, 8, 9, 10, 11, 12, 13)
	printSlice(s)

	printArray(a)
}

func printSlice(s []int){
	fmt.Printf("Len = %d, Cap = %d, Slice = %v\n", len(s), cap(s), s)
}

func printArray(a [5]int){
	fmt.Printf("Len = %d, Cap = %d, Slice = %v\n", len(a), cap(a), a)
}