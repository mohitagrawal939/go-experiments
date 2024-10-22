package main

import "fmt"

func main() {
	var i interface{} =true

	s, ok := i.(int)
	fmt.Println(s, ok)

	s1, ok := i.(string)
	fmt.Println(s1, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}