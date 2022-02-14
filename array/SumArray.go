package main

import (
	"fmt"
	"math/rand"
	"time"
)

func arrSum(arr [10]int) int {
	var Sum = 0
	for _, val := range arr {
		Sum += val
	}
	return Sum
}

func main() {
	// 给定一个随机数种子
	rand.Seed(time.Now().Unix())
	arr := [10]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("数组的内容", arr)
	result := arrSum(arr)
	fmt.Println(result)
}
