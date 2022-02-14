package main

import "fmt"

// 遍历字符串
func traversalString() {
	s := "复习Golang第一天，晚上复习Java"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	fmt.Println("=======================")
	for _, val := range s { // rune
		fmt.Printf("%v(%c)", val, val)
	}
	fmt.Println()
	fmt.Println("=======================")
}

func main() {
	traversalString()
}
