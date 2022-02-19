package main

import "fmt"

func main() {
	allChan := make(chan interface{}, 20)
	for i := 0; i < 20; i++ {
		allChan <- i
	}

	// 如果没有关闭管道，那么会出现deadlock!死锁
	// 如果关闭管道，那么循环也会正常退出，不会出现死锁，管道的关闭只会影响写入，不影响读出
	close(allChan)
	for v := range allChan {
		fmt.Printf("%d\n", v)
	}
}
