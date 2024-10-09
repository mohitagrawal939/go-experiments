package main

import (
	"fmt"
)

func main() {
	fmt.Println("starting pointers")

	i, j := 42,2701;
	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(j)
}
