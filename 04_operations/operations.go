package main

import "fmt"

func sum(a int, b int) int {
	return a+b
}

func multiply(a, b int) int {
	return a*b
}

func substraction(a, b int) int {
	return a-b
}

func main(){
	fmt.Println("Sum of the two number is ", sum(1,2))
	fmt.Println("Multiplication of the two number is ", multiply(1,2))
	fmt.Println("Substraction of the two number is ", substraction(1,2))
}