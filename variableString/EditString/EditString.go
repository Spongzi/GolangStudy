package main

import "fmt"

func changeString() {
	s1 := "hello"
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	s1 = string(byteS1)
	fmt.Println(s1)

	s2 := "我是个程序员!"
	runeS2 := []rune(s2)
	runeS2[0] = '你'
	fmt.Println(string(runeS2))
}
func main() {
	changeString()
}
