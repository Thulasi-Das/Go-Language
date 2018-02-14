package main
import "fmt"
func swap(x, y string) (string, string){
	return y, x
}
func main() {
	a, b := "abc", "bcd"
	fmt.Println("Strings before swapping : ", a, b)
	a, b = swap(a, b)
	fmt.Println("Strings after swapping : ", a, b)
}