package main

import "fmt"

// 字符串格式化
func main() {
	// 打印多个变量
	fmt.Println(1, 2)

	// 万能的%v占位符
	s := "hello"
	n := 123
	arr := [3]int{1, 2}
	fmt.Printf("s=%v\n", s)
	fmt.Printf("s=%v\n", n)
	fmt.Printf("s=%v\n", arr)
	// %+v详细
	// %#v更详细

}
