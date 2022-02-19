package main

import "fmt"

func testArray(arr *[]int) {
	fmt.Printf("func arr: %p %v \n", arr, *arr)
	(*arr)[1] = 100
}

func main() {
	arr1 := [2]int{1, 2} // 传数组
	arr2 := arr1[:]      // 传切片
	testArray(&arr2)
	fmt.Printf("arr1: %p %v \n", &arr1, arr1)
	fmt.Printf("arr2: %p %v \n", &arr2, arr2)
	
}
