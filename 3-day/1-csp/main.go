package main

func main() {
	// 无缓冲通道，同步
	//ch1 := make(chan int)
	// 有缓存通道，解决同步问题
	//ch2 := make(chan int, 2)
	CalSquare()
}

func CalSquare() {
	src := make(chan int)
	dest := make(chan int, 3)

	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()

	for i := range dest {
		println(i)
	}
}

/**
协程之间的通信怎么做？
1. 提倡通过通信共享内存，而不是通过共享内存而实现通信。
	* 通过通信共享内存：channel，通道。
*/
