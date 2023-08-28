package main

import (
	"bufio" // 导入 bufio 包，用于提供带缓冲的 I/O 操作
	"log"   // 导入 log 包，用于记录日志
	"net"   // 导入 net 包，用于进行网络通信
)

func main() {
	// 在本地 1080 端口上监听 TCP 连接
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err) // 如果出现错误，直接 panic 终止程序
	}
	// 循环等待客户端连接
	for {
		client, err := server.Accept() // 接受客户端连接请求
		if err != nil {
			log.Panicln("Accept failed %v", err) // 记录错误日志并继续等待
			continue
		}
		// 启动一个子线程来处理该客户端的请求
		go process(client)
	}
}

// 处理客户端请求的函数
func process(conn net.Conn) {
	defer conn.Close()              // 在函数返回时关闭连接
	reader := bufio.NewReader(conn) // 创建带缓冲的读取器
	for {
		b, err := reader.ReadByte() // 从连接中读取一个字节
		if err != nil {
			break // 如果读取失败，则退出循环
		}
		_, err = conn.Write([]byte{b}) // 将读取到的字节发送回客户端
		if err != nil {
			break // 如果发送失败，则退出循环
		}
	}
}
