package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// 0x是表示16进制，但是传输的时候可以传05即可。
const socks5Ver = 0x05 // SOCKS5协议版本
const cmdBind = 0x01
const atypeIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	err := auth(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
	log.Println("auth success")
}

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ 无需验证
	// X’02’ 用户名和密码验证

	// 客户端：
	// 050100

	// 服务器：
	// 组合即
	// 0500 SOCKS5无需验证
	// 0502 用户名和密码验证

	// 读取一个字节，也就是读取VER
	ver, err := reader.ReadByte()
	// 判断读取的过程中是否有问题
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	// 如果读取到的ver并不是0x05，就拒绝该连接，因为此时连接并不是socks5协议
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	// 读取方法的大小
	methodSize, err := reader.ReadByte()
	fmt.Println(methodSize)
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	// 读取 METHODS 字段
	// 创建了一个长度为methodSize的字节数组
	method := make([]byte, methodSize)
	// 将method全部读到其中
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}
	log.Println("ver", ver, "method", method)
	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+
	// 无需用户名和密码验证

	// curl -x socks5://user:password@127.0.0.1:1080 --proxy-user user:password -v http://www.qq.com

	//_, err = conn.Write([]byte{socks5Ver, 0x03})
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}
	return nil
}
