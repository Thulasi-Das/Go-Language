package main
import "fmt"
func main() {
	i := []int{1, 2, 3}
	fmt.Println(i)

	s := []string{"a", "b"}
	fmt.Println(s)

	b := []bool{true, false, true}
	fmt.Println(b)

	st := []struct{
		i int
		b bool
	}{
		{1, true},
		{2, true},
		{3, false},
		{4, true},
	}
	fmt.Println(st)
}