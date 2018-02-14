package main
import "fmt"
func main() {
	i := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(i)
	a := i[1 : 5]
	fmt.Println(a)
	b := i[:6]
	fmt.Println(b)
	c := i[7:]
	fmt.Println(c)
	d := i[:]
	fmt.Println(d)
	e := a[2:]
	fmt.Println(e)
}