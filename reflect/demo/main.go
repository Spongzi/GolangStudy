package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

// ModifyValue 通过反射来修改值
func ModifyValue(b interface{}) {
	value := reflect.ValueOf(b)
	// 通过Elem来取存在地址中的值
	value.Elem().SetInt(100)
}

func ModifyStudentName(b interface{}) {
	value := reflect.ValueOf(b)
	value.Elem().FieldByName("Name").SetString("张三")
}
func main() {
	num := 10
	fmt.Println("修改前的值是", num)
	ModifyValue(&num)
	fmt.Println("修改后的数组是", num)

	stu := Student{
		Name: "哈哈",
		Age:  10,
	}
	fmt.Printf("地址是%p\n", &stu)
	ModifyStudentName(&stu)
	fmt.Println("修改后的Student内容是", stu)
}
