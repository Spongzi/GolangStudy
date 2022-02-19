package main

import "fmt"

func main() {
	// 演示一个试管
	// 创建一个试管
	intChan := make(chan int, 3)
	// 看看intChan是什么
	fmt.Println(intChan)

	// 向管道写入数据
	intChan <- 10
	intChan <- 100
	intChan <- 200
	// 注意当我们给管道写入数据的时候，数据不可以超出容量
	//intChan <- 2000
	fmt.Printf("看看长度%d,看看容量%d\n", len(intChan), cap(intChan))
	// 从管道中取出数据
	var num2 int
	num2 = <-intChan
	fmt.Println(num2)
	num3 := <-intChan
	fmt.Println(num3)
	mapChan := make(chan map[string]string, 10)
	map1 := make(map[string]string, 20)
	map1["张三"] = "1"
	map1["李四"] = "12"
	map1["王五"] = "13"
	map2 := make(map[string]string, 20)
	map2["李三"] = "2"
	map2["李四"] = "21"
	map2["李五"] = "22"
	mapChan <- map1
	mapChan <- map2
	fmt.Println("mapChan的长度", len(mapChan))
	for i := 0; i < 2; i++ {
		fmt.Println("这是第", i, <-mapChan)
	}
}
