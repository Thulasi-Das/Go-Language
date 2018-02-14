package main
import (
	"fmt"
	"io"
	"os"
	"hash/crc32"
)
func main(){
	h := crc32.NewIEEE()
	w := io.MultiWriter(h, os.Stdout)
	fmt.Fprintf(w, "Hello MultiWriter\n")
	fmt.Printf("hash=%#x\n", h.Sum32())
}