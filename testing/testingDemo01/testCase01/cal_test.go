package testCase01

import (
	"testing"
)

// 编写一个测试用例, 去测试刚刚写的函数是否正确
func TestAddUpper(t *testing.T) {
	// 调用函数
	res := addUpper(10)
	if res != 55 {
		//fmt.Println()
		t.Fatalf("执行错误，期望值是55， 实际值是%v", res)
	}
	//如果正确，输入日志
	t.Logf("执行正确")
}
