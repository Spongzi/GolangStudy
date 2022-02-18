package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

// JSON 格式化数据
func testStruct() {
	// 首先先声明创建一个结构体
	monster := Monster{
		Name:     "牛魔王",
		Age:      20,
		Birthday: "2002-6-15",
		Sal:      1000,
		Skill:    "牛魔冲撞",
	}
	// 序列化
	bytes, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}

func testMap() {
	var map1 = map[string]interface{}{
		"name": "苏默认",
		"age":  20,
	}
	bytes, err := json.Marshal(map1)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	fmt.Println(string(bytes))
}
func testSlice() {
	var slice = []map[string]interface{}{
		{"name": "su", "age": 18},
		{"name": "wang", "age": 20},
		{"name": "wang1", "age": 21},
		{"name": "wang2", "age": 22},
	}
	bytes, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("Json序列化失败, err", err)
		return
	}
	fmt.Println(string(bytes))
}
func main() {
	testStruct()
	testMap()
	testSlice()
}
