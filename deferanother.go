package main
import "fmt"
func main() {
	fmt.Println("Start")
	defer fmt.Println("Defer Statement")
	fmt.Println("End")
}