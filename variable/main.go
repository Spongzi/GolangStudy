package main

import "fmt"

// 全局变量
var m = 100

func main() {
	n := 10
	m := 200 // 此时声明变量为m
	fmt.Println(n, m)
}
