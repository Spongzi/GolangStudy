package main

import (
	"fmt"
	"time"
)

// 计算1-200的各个数的阶乘，并将其放入map中
// 定义一个函数，来求阶乘，并且放入到map中
// 启动多个协程，统一的将结果放入map中,map应该是全局的

var (
	myMap = make(map[int]int, 10)
)

func test(n int) {
	res := 1
	for i := 0; i < n; i++ {
		res *= i
	}
	myMap[n] = res
}

func main() {
	now := time.Now()
	for i := 0; i < 200; i++ {
		go test(i)
	}
	// 输出结果
	for k, v := range myMap {
		fmt.Printf("%v! = %v", k, v)
	}
	end := time.Now()
	fmt.Println("一共运行了", end.Sub(now))
}

// 271
