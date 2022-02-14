package main

import "fmt"

func ptr(num *int) {
	*num = 100
}
func main() {
	var num int
	num = 10
	fmt.Printf("num: %p\n", &num)
	ptr(&num)
	fmt.Println(num)
}
