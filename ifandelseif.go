package main
import (
	"fmt"
	"math"
)
func pow(x, y, limit float64) float64 {
	if v:= math.Pow(x, y); v < limit{
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit
}
func main() {
	fmt.Printf("3 to the power 2 is(if within limit 10) : %g\n", pow(3,2,10))
	fmt.Printf("3 to the power 4 is(if within limit 10) : %g\n", pow(3,4,10))
}