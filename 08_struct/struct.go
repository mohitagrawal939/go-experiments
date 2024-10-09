package main

import (
	"fmt"
)

type Vertex struct {
	x int
	Y int
}
func main() {
	fmt.Println(Vertex{1,2})
	v := Vertex{2,3}
	fmt.Printf("%T",v)
}
