package myCalculator

import "fmt"

func Plus(A int, B int) int{
	return A+B
}

func Minus(A int, B int) int{
	return A-B
}

func Out(m string) {
	fmt.Printf("text: %s", m)
}