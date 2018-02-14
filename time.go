package main
import (
	"fmt"
	"time"
)
func main(){
	day := time.Now().Weekday()
	start := time.Now()
	fmt.Printf("Day of the Week : %s (%d)\n", day, day)
	fmt.Printf("%s\n", time.Now())
	fmt.Println(time.Since(start))
	fmt.Println(time.Hour + time.Since(start))
}