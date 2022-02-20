package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

// 案例 对进本数据类型进行一个控制
func test01(b interface{}) {
	// 获取type, kind, 值
	bTyp := reflect.TypeOf(b)
	fmt.Println("获得的Type = ", bTyp)
	// 获取值
	bVal := reflect.ValueOf(b)
	n1 := 2 + bVal.Int()
	fmt.Println("得到n1 = ", n1)
	fmt.Println("获取的Value", bVal)
	// 将类型转化为interface
	iv := bVal.Interface()
	// 将interface通过断言重新转成需要的类型
	num := iv.(int)
	fmt.Printf("num = %d, num type = %T\n", num, num)
}

func test02(b interface{}) {
	// 获取type, kind, 值
	bTyp := reflect.TypeOf(b)
	fmt.Println("获得的Type = ", bTyp)
	// 获取值
	bVal := reflect.ValueOf(b)
	fmt.Println("获取的Value", bVal)
	// 将类型转化为interface
	bi := bVal.Interface()
	// 类型断言两种方式
	// 可以使用多一个变量接受bool来判断
	// 也可以使用switch 的type来判断
	stu, ok := bi.(Student)
	if ok {
		fmt.Println("学生的名字=", stu.Name)
	} else {
		fmt.Println("我不知道你是什么东西，滚开")
	}
}
func main() {
	// 先定义一个int类型
	var num = 100
	test01(num)
	// 定义一个Student的实例
	stu := Student{Name: "张三", Age: 10}
	test02(stu)
}
