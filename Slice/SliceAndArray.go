package main

import "fmt"

func main() {
	var arr1 = [2]int{1, 2}
	var arr2 [2]int
	arr2 = arr1
	fmt.Printf("arr1: %p %v \n", &arr1, arr1)
	fmt.Printf("arr2: %p %v \n", &arr2, arr2)
}
