package main
import (
	"fmt"
	"encoding/json"
	"log"
)
type Lang struct{
	Name string
	year int
	URL string
}
func main() {
	lang := Lang{"Go", 2009, "http://golang.org"}
	data, err := json.MarshalIndent(lang, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}