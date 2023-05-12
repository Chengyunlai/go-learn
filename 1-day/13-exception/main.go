package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(isEquals("张三", "李四")) //false 不等
	fmt.Println(isEquals("张三", "张三")) //false 不等

	value, err := isEquals("张三", "李四")
	if err != nil {
		// 打印错误消息
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

// 错误是error类型，通过errors.xx 创建相应的错误
func isEquals(name, input string) (bool, error) {
	if name != input {
		return false, errors.New("不等")
	}
	// nil 表示 null
	return true, nil
}
