package main

import "testing"

func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "haha",
		Age:   18,
		Skill: "三昧真火",
	}
	res := monster.Store()
	if res {
		t.Logf("正确，一点毛病没有")
	} else {
		t.Fatalf("错误，你赶快去看看哪里报错了，我程序写的不好，我应该返回err，但是我返回了true和false")
	}
}
