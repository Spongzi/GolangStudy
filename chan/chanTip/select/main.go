package main

import "fmt"

func main() {
	// 使用select可以解决从管道取数据的阻塞问题
	// 1. 定义一个管道10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// 2. 定义一个管道5个数据string
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		str := fmt.Sprintf("stu%d", i)
		strChan <- str
	}
	// 传统方法，如果不关闭管道遍历管道的话会产生deadlock
	// select也可以解决
	for true {
		select {
		case v := <-intChan:
			// 注意这里如果管道最终都没有关闭，也不会一直阻塞而导致死锁(deadlock)
			// 会自动的到下一个case匹配
			fmt.Println("从intChan中读取了数据", v)
		case v := <-strChan:
			fmt.Println("从strChan中读取了数据", v)
		default:
			fmt.Println("没有数据了，程序员可以加入业务逻辑")
			return
		}
	}
}
