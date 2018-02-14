package main
import (
	"fmt"
	"encoding/xml"
	"log"
)
type Lang struct{
	Name string
	year int
	URL string
}
func main() {
	lang := Lang{"Go", 2009, "http://golang.org"}
	data, err := xml.MarshalIndent(lang, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}