package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// 定义一个存放任意数据类型的管道3个数据
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "Tom"
	allChan <- Cat{
		Name: "嘻嘻",
		Age:  19,
	}
	// 我们希望获得管道中的第三个元素，则现将前两个退出
	<-allChan
	<-allChan
	data := <-allChan
	fmt.Printf("data的数据类型是%T", data)
	// 使用类型断言可以使用字段
	// 如果直接使用data访问字段是不可以的！！！
	newCat := data.(Cat)
	fmt.Println(newCat.Name)
}
