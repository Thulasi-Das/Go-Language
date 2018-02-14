package main
import (
	"fmt"
)
func main() {
	fmt.Println("counting")
	for i := 0; i <= 10; i++{
		//prints in reverse order
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}