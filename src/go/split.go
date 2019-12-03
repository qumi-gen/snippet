package main
import (
	"strings"
	"fmt"
)
func main() {
	src := "ABCD DEF GHI JKL"
	arr := strings.Split(src, " ")
	fmt.Printf("%s\n", arr[0])
}
