package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 100
	println(a[0], len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b) // 得用这个才能直接打印出数组

	// 二维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

}
