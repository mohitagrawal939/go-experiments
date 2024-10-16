package main

import (
	"fmt"
)

func main() {
	a := [...]int{1,2,3,4,5}
	a = [...]int{1, 3:100, 5} // index 2,3 will have zero values
	fmt.Println(a)
}
