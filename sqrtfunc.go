package main
import (
	"fmt"
	"math"
)
func Sqrt(x float64) float64 {
	var z float64 = 1.0
	i := 0
	const prec = 5
	for z < x {
		fmt.Printf("value of z = %.100f\n", z)
		//prev := z
		z -= ((z * z) - x) / (2 * z)
		i++
		fmt.Printf("z after %v iteration is %.100f\n", i, z)
		// if i == 10 {
		// 	fmt.Println("breaking after 10 iterations")
		// 	break
		// }
		// if float32(prev) == float32(z) {
		// 	fmt.Println("breaking due to not much change in digits")
		// 	break
		// }
		// temp := float32(x)
		// fmt.Printf("temp : %f\n", temp)
		// fmt.Printf("y = x/2 : %v\n", float32(temp/2))
		// fmt.Printf("float32(z) : %v\n", float32(z))
		// if y := x/2; z <= y {
		// 	fmt.Println("breaking as z is less than or equal to x/2")
		// 	break
		// }
		// if i == int(x) {
		// 	fmt.Printf("breaking after %v iterations", x)
		// 	break
		// }
		if j := x/2; i == int(j) {
			fmt.Printf("breaking after %v iterations(that is the half of %v)", j, x)
			break
		}
		fmt.Println();
	}
	return z
}
func main() {
	//fmt.Printf("Sqrt of %v is %g\n", 2, Sqrt(2))
	fmt.Printf("Actual Sqrt of %v with math.Sqrt func is %v",2, math.Sqrt(2))
}