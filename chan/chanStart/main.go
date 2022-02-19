package main

import (
	"fmt"
	"sync"
	"time"
)

// 计算1-200的各个数的阶乘，并将其放入map中
// 定义一个函数，来求阶乘，并且放入到map中
// 启动多个协程，统一的将结果放入map中,map应该是全局的

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁
	// sync 是包：synchronized 同步锁
	// Mutex 是互斥锁
	lock sync.Mutex
	wg   sync.WaitGroup
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 执行之前，加锁
	lock.Lock()
	myMap[n] = res
	// 执行结束后要解锁
	defer lock.Unlock()
}

func main() {
	now := time.Now()
	for i := 1; i < 20; i++ {
		go test(i)
	}
	// 输出结果
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("%v! = %v\n", k, v)
	}
	lock.Unlock()
	end := time.Now()
	fmt.Println("一共运行了", end.Sub(now))
}

// 271
