package main

import "fmt"

func main() {
	// key string, value int
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(len(m))     //2
	fmt.Println(m["one"])   //1
	fmt.Println(m["two"])   //2
	fmt.Println(m["three"]) //0

	// ok 获取map中是否有这个key存在
	r, ok := m["three"]
	//fmt.Println(r)
	//fmt.Println(ok)
	fmt.Println(r, ok)

	// 删除key为one
	delete(m, "one")

	// 简易定义
	m2 := map[string]int{"one": 1, "two": 2}
	fmt.Println(m2)

	// 遍历map，同理可以遍历数组
	for key, value := range m {
		println(key, value) // key,value
	}
}

/**
map 是无序的
*/
