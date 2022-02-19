package main

import "fmt"

func main() {
	// 管道可以直接定义为只读或者只写
	// 默认情况下管道是双向的，即可以读也可以写
	// var chan1 chan int
	// 只写
	var chan2 chan<- int
	chan2 = make(chan int, 10)
	chan2 <- 1
	// 这样就会报错
	// num := <-chan2
	fmt.Println(chan2)
	// 只读
	var chan3 <-chan int
	chan3 = make(chan int, 10)
	// 报错
	// chan3 <- 10
	num := <-chan3
	fmt.Println(num)
}
