package main

import (
	"fmt"
)

func main() {
	fmt.Println("testing");

	defer fmt.Println("second print due to defer")

	fmt.Println("First print")
}
