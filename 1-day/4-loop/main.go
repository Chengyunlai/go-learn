package main

func main() {
	// 死循环=while(true)
	for {
		println(1)
		break
	}

	// 打印0-9
	for i := 0; i < 10; i++ {
		println(i)
	}
}

/**
go没有while，没有do while，只有for循环
*/
