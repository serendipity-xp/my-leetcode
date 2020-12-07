package str

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	a := []int32("abc")
	b := []int32("a23")
	fmt.Println(a[0])
	fmt.Println(len("abc"), len(a))
	fmt.Println(string(a[0]+b[0]))
}

func TestStrOperation(t *testing.T) {
	binary := addBinary("1001", "0101", 2)
	fmt.Println(binary)
}