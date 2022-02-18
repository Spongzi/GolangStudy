package main

import (
	"fmt"
	"strconv"
	"time"
)

// goroutine快速入门

// 写一个函数每隔一秒钟输出一次HelloWorld
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello, World" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("Hello golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
