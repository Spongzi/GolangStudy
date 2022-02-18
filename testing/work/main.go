package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Monster 编写一个结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

// Store 给Monster添加一个方法可以将对象序列化之后保存到文件
func (m *Monster) Store() bool {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json failed", err)
		return false
	}
	file, err := os.OpenFile("./JSON数据.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed", err)
		return false
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("close failed", err)
			return
		}
	}(file)
	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		fmt.Println("writer failed", err)
		return false
	}
	_ = writer.Flush()
	return true
}

func (m *Monster) ReStore() bool {
	file, err := os.Open("./JSON数据.json")
	if err != nil {
		fmt.Println("file open failed", err)
		return false
	}
	reader := bufio.NewReader(file)
	var fileContent string
	for true {
		readRune, _, err := reader.ReadRune()
		if err == io.EOF {
			fmt.Println("read file failed", err)
			break
		}
		fileContent += string(readRune)
	}
	fmt.Println(fileContent)
	return true
}

func main() {
	m := Monster{
		Name:  "su",
		Age:   19,
		Skill: "吃喝玩乐",
	}
	m.ReStore()
}
