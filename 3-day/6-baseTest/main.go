package main

import (
	"github.com/bytedance/gopkg/lang/fastrand"
	"math/rand"
)

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}

// 对Select函数做基本测试
func Select() int {
	return ServerIndex[rand.Intn(10)]
}

func FastSelect() int {
	return ServerIndex[fastrand.Intn(10)]
}
