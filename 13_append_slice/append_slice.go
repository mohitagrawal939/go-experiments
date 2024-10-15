package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Printf("s len=%d cap=%d %v \n", len(s), cap(s), s)

	s = append(s, 0)
	fmt.Printf("s len=%d cap=%d %v \n", len(s), cap(s), s)

	s = append(s, 0)
	fmt.Printf("s len=%d cap=%d %v \n", len(s), cap(s), s)
}
