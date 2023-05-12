package main

import (
	"net/http"
)

func main() {
	println("服务器启动")
	// 处理http的/请求，映射到服务器的文件系统路径："W:\\"，进入网页中可以到自己主机W盘的内容
	http.Handle("/", http.FileServer(http.Dir("W:\\")))
	// 服务器启动，以localhost:8080访问
	http.ListenAndServe("localhost:8080", nil)
}

/**
初探Go的魅力，访问localhost:8080，一个完整的静态资源ftp服务器展现在眼前
*/
