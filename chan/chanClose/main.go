package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 20
	intChan <- 19
	close(intChan) // close 内置函数，专门用来关闭管道
	// 这时候就不能够再写入数据到channel
	// intChan <- 200
	// 但是这个时候可以读
	fmt.Println(<-intChan)
}
