package main

import "fmt"

func main() {
	// 用make创建切片，需要指定初始长度
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3

	// 新增，需要用一个切片变量接收，容量不够会扩容，返回一个新的slice
	s = append(s, "d")
	sNew := append(s, "e", "f")
	fmt.Println(sNew)

	// copy
	a := make([]string, len(sNew))
	copy(a, sNew)
	fmt.Println(a)

	// 切片操作
	fmt.Println(a[:])   // 所有
	fmt.Println(a[2:5]) // 区间[2,5)
	fmt.Println(s[:3])  // [0,3)
	fmt.Println(s[1:])  // 1到结尾

	// 简易定义
	good := []string{"g", "o", "o", "d"}
	fmt.Println(good)
	good = append(good, "!")
	fmt.Println(good)
	// 区别数组：	b := [5]int{1, 2, 3, 4, 5}，是需要指定长度的
}

/**
切片是一个可变长数组，不用指定其长度。[]类型
*/
