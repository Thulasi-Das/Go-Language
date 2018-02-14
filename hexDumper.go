package main
import (
	"encoding/hex"
	"fmt"
	"os"
)
func main(){
	hexDump := hex.Dumper(os.Stdout)
	defer hexDump.Close()
	fmt.Fprintf(hexDump, "Hello, hexDumper")
}