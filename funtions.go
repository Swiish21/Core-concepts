package main

import "fmt"

var addition = func (a int, b int) int {
	return a + b
}

var subtraction = func (c int, d int) {
	return b - c
} 

func main() {
	fmt.Println(addition(1, 2))
}