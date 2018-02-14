package main
import "fmt"
func c(i int) int {
	defer func() {
		i++
		fmt.Println("Inside defer function ", i)
	} ()
	//i++
	return i
}
func main() {
	fmt.Println("Value returned : ", c(1))
}