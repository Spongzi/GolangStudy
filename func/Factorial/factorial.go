package main

import "fmt"

// TODO 阶乘
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
func main() {
	fmt.Println(factorial(10))
}
