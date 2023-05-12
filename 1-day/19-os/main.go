package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// go run 这个程序 a b c d
	fmt.Println(os.Args) // [C:\Users\12579\AppData\Local\Temp\go-build3318116271\b001\exe\main.exe a b c d]

	fmt.Println(os.Getenv("PATH")) // 获取环境变量
	//fmt.Println(os.Setenv("AA","BB")) // 写入环境变量

	println("----")
	buf, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}

/**
进程信息，目前不知道有什么作用
*/
