package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建一个Person结构体[Name, Age, Address]
// 使用rand方法配合随创建10个Person实例，放入到channel中
// 遍历channel，将各个Person的信息显式在终端
type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	personChan := make(chan Person, 10)
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("王%d", rand.Intn(100))
		address := fmt.Sprintf("南山路第%d号街", rand.Intn(100))
		personChan <- Person{
			Name:    name,
			Age:     rand.Intn(100),
			Address: address,
		}
	}
	lenChan := len(personChan)
	for i := 0; i < lenChan; i++ {
		fmt.Println(<-personChan)
	}
}
