package main

func main() {
	n := 5
	add(n)
	println(n) // 5
	add2(&n)
	println(n) // 7
}

// 两个函数均没有返回值
func add(n int) {
	// 此时的n是add的局部变量，修改完后，当add函数执行完毕后，n消亡，不影响外部的变量值
	n += 2
}

func add2(n *int) {
	// 此时的n是一个指针，改动会改变原值
	*n += 2
}

/**
指针的主要用途，对传入的参数进行修改
*/
