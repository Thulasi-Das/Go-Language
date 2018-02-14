package main
import (
	"os"
	"io"
	"log"
)
func main() {
	input, err := os.Open("testjson.json")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, input)
} 