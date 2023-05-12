package main

import "fmt"

func main() {
	arr := [3]int{}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	// index是数组下标，value对应的值
	for index, value := range arr {
		fmt.Println(index, value)
	}

	m := map[string]int{"one": 1, "two": 2}
	
	//key是map的key，value是map的value
	for key, value := range (m) {
		fmt.Println(key, value)
	}
}
