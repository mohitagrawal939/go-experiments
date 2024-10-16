package main

import (
	"fmt"
)

func main() {
	source := []int{2,3,5,7,11}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Println(source)
	fmt.Println(destination)
}
