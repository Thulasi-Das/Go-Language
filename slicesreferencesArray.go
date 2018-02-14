package main
import "fmt"
func printString(s [5]string){
	fmt.Println(s)
}
func printTwoStrings(s1, s2 []string){
	fmt.Println(s1, s2)
}
func main() {
	names := [5]string{"Tulc", "Daz", "Duggu", "Guddu", "Dude"}
	s1 := names[0 : 2]
	s2 := names[1 : 3]
	printString(names)
	printTwoStrings(s1, s2)

	s1[0] = "Thulasi"
	s2[0] = "Das"

	printString(names)
	printTwoStrings(s1, s2)
}