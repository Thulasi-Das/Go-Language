package main

import "fmt"

func main(){
	m := make(map[string]int)

	m["One"] = 1
	fmt.Println(m["One"])

	m["One"] = 11
	fmt.Println(m["One"])

	delete(m, "One")
	fmt.Println(m["One"])

	elem, ok := m["One"]
	fmt.Printf("Element : %d\tPresent : %v\n", elem, ok)
}