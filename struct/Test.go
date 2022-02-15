package main

import "fmt"

type student1 struct {
	id   int
	name string
	age  int
}

func demo(ce []student1) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	ce = append(ce, student1{3, "xiaowang", 56})
	// return ce
}
func main() {
	var ce []student1 //定义一个切片类型的结构体
	ce = []student1{
		student1{1, "xiaoming", 22},
		student1{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
}
