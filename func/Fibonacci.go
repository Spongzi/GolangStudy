package main

import "fmt"

// Fibonacci 斐波那契数列
func Fibonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return Fibonacci(i-1) + Fibonacci(i-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Fibonacci(i))
	}
}
