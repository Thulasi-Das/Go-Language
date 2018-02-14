package main
import (
	"fmt"
	"io"
	"os"
)
func copyStr(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Open(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(src, dst)
}
func main() {
	returnStr, err := copyStr("src.txt", "dst.txt")
	fmt.Println(returnStr, err)
}