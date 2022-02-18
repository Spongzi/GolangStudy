package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取的CPU个数
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
	// 可以设置使用多少个CPU
	// 1.8以后的版本默认是在多核上运行
	runtime.GOMAXPROCS(numCPU - 1)
	fmt.Println("ok")
}
