package main
import (
	"fmt"
	"time"
)
func main() {
	fmt.Println("Time now : ", time.Now())
	fmt.Println(" Saturday : ", time.Saturday)
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	fmt.Println("Today : ", today)
	switch time.Saturday {
		case today + 0 :
			fmt.Println("Today")
		case today + 1 :
			fmt.Println("Tomorrow")
		case today + 2 : 
			fmt.Println("Two days from now")
		case today + 3 :
			fmt.Println("Three days from now")
		default :
			fmt.Println("Too far ahead")
	}
}