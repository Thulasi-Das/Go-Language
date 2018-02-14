package main
import "fmt"
func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}

	s:= a[0 : 3]
	printSlice(s)

	s = s[ : 2]
	printSlice(s)

	s = s[1 : ]
	printSlice(s)

	s = s[2 : 5]
	printSlice(s)

	//panics as the higher bound is greater than the capacity of the previously sliced s
	// s = s[0 : 4]
	// printSlice(s)
}
func printSlice(s []int){
	fmt.Printf("Len = %d, Cap = %d, Slice = %v\n", len(s), cap(s), s)
}