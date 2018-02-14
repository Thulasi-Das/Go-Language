package main

import "fmt"

func main() {
	var s []int
	if s == nil {
		fmt.Println("Nil slice")
		fmt.Println(s)
	}
}