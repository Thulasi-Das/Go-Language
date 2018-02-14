package main

import (
	//"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var words = strings.Fields(s)
	var m = make(map[string]int)
	//fmt.Printf("%q", words)
	for _, val := range words {
		_, ok := m[val]
		if ok {
			m[val] += 1
		} else {
			m[val] = 1
		}
	}
	fmt.Println(m)
	return m
}

func main() {
	WordCount("This is a test String String")
	//wc.Test(WordCount)
}
