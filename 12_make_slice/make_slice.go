package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	fmt.Printf("a len=%d cap=%d %v \n", len(a), cap(a), a)

	b := make([]int, 0, 5)
	fmt.Printf("b len=%d cap=%d %v \n", len(b), cap(b), b)
}
