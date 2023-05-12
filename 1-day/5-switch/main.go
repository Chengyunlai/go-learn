package main

import "time"

func main() {
	a := 2
	switch a {
	case 1:
		println("one")
	case 2:
		println("two")
	default:
		println("不知道")
	}

	t := time.Now()
	println(t)
	// 特殊用法，取代if - else
	switch {
	case t.Hour() < 12:
		println("早于中午")
	default:
		println("晚于中午")
	}
}

/**
switch自带一个break
*/
