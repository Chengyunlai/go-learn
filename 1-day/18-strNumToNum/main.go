package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 一个10进制字符串数字，转到X进制，精度是64位
	f, _ := strconv.ParseInt("111", 2, 64)
	fmt.Println(f) // 7

	n, _ := strconv.Atoi("123")
	fmt.Println(n) // 转数字
}
