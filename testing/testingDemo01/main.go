package main

import "fmt"

func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	// 传统的测试方法，就是在main方法中运行查看结果是否正确
	res := addUpper(10) // 55
	if res != 55 {
		fmt.Println("addUpper错误, 返回值", res, "期望值是55")
	} else {
		fmt.Println("正确")
	}
	/*
		缺点：
		1. 需要修改主函数的代码，如果项目在运行，就可能需要去停止项目
		2. 不利于管理，因为当我们需要测试多个函数或者多个模块的时候，都需要写在main函数中，不利于我们管理和清晰我们的思路4
		3. 引出单元测试 --> testing
	*/
}
