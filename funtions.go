package main

import "fmt"

var addition = func (a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(addition(1, 2))
}