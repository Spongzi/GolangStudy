package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello world")
	}
}

func Test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test发生了错误", err)
		}
	}()
	// 定义了一个map
	var myMap map[int]string
	myMap[0] = "Golang" // err
}
func main() {
	go sayHello()
	go Test()
	for i := 0; i < 10; i++ {
		fmt.Println("main", i)
	}
}
