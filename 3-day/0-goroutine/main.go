package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	WaitHelloGoRoutine()
}

func hello(i int) {
	println("helo goroutine:" + fmt.Sprint(i))
}

func HelloGoRoutine() {
	for i := 0; i < 5; i++ {
		// go 关键字，就是开启协程，例如下面这个将这个方法交给协程去运行
		go func(j int) {
			hello(j)
		}(i)
	}
	// 会有更优雅的方式
	time.Sleep(time.Second)
}

func WaitHelloGoRoutine() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// go 关键字，就是开启协程，例如下面这个将这个方法交给协程去运行
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	// 会有更优雅的方式
	//time.Sleep(time.Second)
	wg.Wait()
}

/**
如何开启一个协程
*/
