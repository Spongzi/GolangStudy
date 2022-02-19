package main

import "fmt"

/*
	FindNum
 	找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，找出两个元素之和等于8的下标分别是（0，4）和（1，2）
*/
func FindNum(arr [5]int, target int) {
	for i := 0; i < len(arr); i++ {
		other := target - arr[i]
		for j := i; j < len(arr); j++ {
			if other == arr[j] {
				fmt.Printf("(%d, %d)", i, j)
			}
		}
	}
}
func main() {
	arr := [5]int{1, 3, 5, 8, 7}
	FindNum(arr, 8)
}
