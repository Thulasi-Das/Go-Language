package main
import "fmt"
var packageSum int
func split(sum int)(x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	sum := 40
	var functionSum int
	functionSum = 35
	packageSum = 30
	fmt.Println(split(sum))
	fmt.Println("package sum split")
	fmt.Println(split(packageSum))
	fmt.Println("function sum split")
	fmt.Println(split(functionSum))
}