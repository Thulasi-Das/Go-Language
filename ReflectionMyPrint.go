package main
import (
	"reflect"
	"os"
	"strconv"
)
func main() {
	myPrint("Hello World\t", 24, "\tage")
}
func myPrint(args... interface{}){
	for _, arg := range args {
		switch v:= reflect.ValueOf(arg); v.Kind(){
			case reflect.String : 
				os.Stdout.WriteString(v.String())
			case reflect.Int : 
				os.Stdout.WriteString(strconv.FormatInt(v.Int(), 10))
		}
	}
}