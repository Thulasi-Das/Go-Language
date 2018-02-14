package main
import "fmt"
func main() {
	i,j := 21, 2701
	p := &i
	fmt.Println(*p)
	*p = 22
	fmt.Println(i)

	p = &j
	fmt.Println(*p)
	*p = *p / 37
	fmt.Println(j)
}