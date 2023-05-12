package main

func main() {
	println(add(1, 2))
	println(exists(map[string]int{"one": 1}, "two"))
}

// 函数名add，参数a,b均为int，返回值为int
func add(a, b int) int {
	return a + b
}

// 支持多值返回
func exists(m map[string]int, key string) (int, bool) {
	value, ok := m[key]
	return value, ok
}
