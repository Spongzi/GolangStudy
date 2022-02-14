package main

import "fmt"

func main() {
	// map的遍历
	scoreMap := map[string]int{
		"张三": 100,
		"王五": 110,
	}
	for k, v := range scoreMap {
		fmt.Println("key:", k, "value", v)
	}
	for k := range scoreMap {
		fmt.Println(k)
	}
}
