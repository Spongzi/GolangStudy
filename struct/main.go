package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "张三", age: 17},
		{name: "李四", age: 19},
		{name: "王五", age: 29},
	}
	for _, stu := range stus {
		fmt.Printf("stu: %v %p\n", stu, &stu)
		m[stu.name] = &stu
	}
	for key, value := range m {
		fmt.Println(key, "=>", value.name)
	}
	var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	for i, ints := range arr1 {
		fmt.Println(i, ints)
	}
	fmt.Println(len(arr1))
}
