package main

//import "golang.org/x/tour/pic"
import "fmt"

func printSliceMatrix(s [][]uint8) {
	for _, v := range s {
		fmt.Printf("%d", v)
	}
	fmt.Println()
}

func Pic(dx, dy int) [][]uint8 {
	var p = make([]([]uint8), dy)
    for i := 0; i < len(p); i++ {
        p[i] = make([]uint8, dx)
        for j := 0; j < len(p[i]); j++ {
        	var value int
        	if i % 2 == 0 {
				value = (i + j)/2
			} else if i % 3 == 0 {
				value = i * j
			} else if i % 1 == 0 {
				value = i ^ j
			}
			p[i][j] = uint8(value)

        }
    }
    printSliceMatrix(p)
    return p
}

func main() {
	//pic.Show(Pic)
	Pic(8, 8)
}
