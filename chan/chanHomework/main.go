package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 开启一个writeData协程，向管道中写入intChan 50个整数
// 开启一个readData协程，从管道中读出写入的数据
// 注意：writeData和readData操作的是同一个管道
// 主线程需要等待writeData和readData都完成后才退出

// 思路：
// 启动两个协程，第一个携程为writeData,另一个为readData
// 对主线程处理，防止writeData和readData没有执行结束的时候，主线程结束
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		// 放入数据
		intChan <- i
	}
	close(intChan) // 关闭
}

func readData(intChan chan int) {
	defer wg.Done()
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("读取到的数据是%v\n", v)
	}
}

func main() {
	intChan := make(chan int, 50)
	wg.Add(1)
	go writeData(intChan)
	wg.Add(1)
	go readData(intChan)
	wg.Wait()
}
