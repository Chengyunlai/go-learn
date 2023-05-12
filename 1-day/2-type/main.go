package main

func main() {
	// 变量声明的方式，以整型为例子
	var intNum1 int = 0
	var intNum2, intNum3 int = 1, 2
	var intNum4 = 3
	// 简易声明，类型自动推断
	intNum5 := 5

	println(intNum1)
	println(intNum2, intNum3)
	println(intNum4)
	println(intNum5)

	// 常量
	const PI float32 = 3.1415926
	println(PI)

	// 布尔
	flag := false
	println(flag)

	// 浮点型64
	var float64Num float64
	println(float64Num)

	// 浮点型32
	float32Num := float32(float64Num)
	println(float32Num)

	// 字符串
	str := "你好"
	str += "世界"
	println(str)
}

/**

Go是一门强类型的语言
*/
