package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 求素数使用goroutine和channel

func putNum(intChan chan int) {
	defer wg.Done()
	for i := 1; i < 80000; i++ {
		intChan <- i
	}
	// 关闭
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int) {
	defer wg.Done()
	var flag bool
	// 使用for循环
	for true {
		num, ok := <-intChan
		if !ok {
			break
		}
		// 判断num是不是一个素数
		flag = isPrimeNum(num)
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("我完成了")
}

func isPrimeNum(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	intChan := make(chan int, 10000)
	// 放入结果
	primeChan := make(chan int, 100000)
	start := time.Now()
	// 开启一个协程, 存入数据
	wg.Add(1)
	go putNum(intChan)
	// 开启四个协程，取出数据，并且判断是否为素数
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan)
	}
	wg.Wait()
	close(primeChan)
	//for i := range primeChan {
	//	fmt.Println(i)
	//}
	end := time.Now()
	timeSub := end.Sub(start)
	fmt.Println("消耗的时间", timeSub)
}
