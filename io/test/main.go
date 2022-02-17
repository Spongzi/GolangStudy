package main

// flag 包的运用！

import (
	"flag"
	"fmt"
)

func main() {
	// 定义几个参数变量
	var user string
	var password string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&password, "pwd", "", "密码默认为空")
	flag.StringVar(&host, "h", "", "链接地址，默认为空")
	flag.IntVar(&port, "p", 1, "端口号，默认为1")
	flag.Parse()
	fmt.Printf("user=%v, password=%v, host=%v, port=%v", user, password, host, port)
}
