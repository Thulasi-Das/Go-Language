package main
import (
	"fmt"
	"math/rand"
)
func main() {
	rand.Seed(2)
	fmt.Println("Random number generated is  : ", rand.Intn(10))	
}