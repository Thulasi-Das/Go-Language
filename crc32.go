package main
import (
	"fmt"
	"hash/crc32"
)
func main(){
	h := crc32.NewIEEE()
	fmt.Fprintf(h, "Hello hash crc32")
	fmt.Printf("hash = %#x\n", h.Sum32) //check sum value
}