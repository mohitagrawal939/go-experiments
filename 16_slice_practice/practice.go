package main

import (
	"fmt"
)

func main() {
	sl := make([]int, 5)
	sl[0] = 10
	fmt.Println(sl)
	one(sl)
	fmt.Println(sl)
	two(sl)
	fmt.Println(sl)
}

// here newSlice is noting but pointing towards sl slice because slice is noting but glorified pointer it doesnt hold value
// it holds the pointing location
// 
func one(newSlice []int){
	newSlice[4] = 11
}

func two(sl []int){
	sl = append(sl, 12)
}
