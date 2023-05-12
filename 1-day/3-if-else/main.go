package main

func main() {
	// if - else
	if 7%2 == 0 {
		// 偶数
		println("7 is even")
	} else {
		// 奇数
		println("7 is odd")
	}

	// if - else if - else
	// num给其一个初始条件
	if num := 9; num < 0 { // 如果 num < 0
		println(num, "is negative") // 输出它是一个负数
	} else if num == 0 {
		println(num, "is zero") // 输出它是一个0
	} else {
		println(num, "is positive") // 输出它是一个正数
	}
}
