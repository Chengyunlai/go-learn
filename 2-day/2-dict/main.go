package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io" // 处理I/O的标准库
	"log"
	"net/http"
	"os"
	"strings"
)

// 定义请求的结构体，我们可以手动输入数据并且封装该请求体为Json的方式传递
type DictRequest struct {
	//{"trans_type":"zh2en","source":"你好"}
	TransType string `json:"trans_type"`
	Source    string `json:"source""`
}
type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func main() {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建一个字符串类型的数据，存放JSON格式的请求数据
	//var data = strings.NewReader(`{"trans_type":"zh2en","source":"你好"}`)

	// 创建对象（改造其内容）
	//request := DictRequest{TransType: "en2zh", Source: "happy"}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请仔细阅读：该程序是【英翻中】词典程序，输入exit-q时退出程序。")
	for {
		// 输入内容
		inputStr := inputChinese(reader)
		inputStr = strings.TrimSuffix(inputStr, "\r\n")
		if inputStr == "exit-q" {
			fmt.Println("欢迎你的使用")
			break
		}
		// 封装请求体
		request := DictRequest{TransType: "en2zh", Source: inputStr}
		// 将其转为json格式
		// 这里buf是一个byte数组，所以下面需要用bytes.NewReader去做转换。
		buf, err := json.Marshal(request)
		if err != nil {
			log.Fatal(err)
		}
		var data = bytes.NewReader(buf)
		// 创建一个新的POST请求
		req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)

		if err != nil {
			log.Fatal(err)
		}

		// 设置请求头
		setHeader(req)

		// 发起请求
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		// 在函数结束后关闭HTTP响应的资源流
		defer resp.Body.Close()

		// 读取响应内容
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// 打印响应内容，得到请求的结果。
		//fmt.Printf("%s\n", bodyText)
		// {"rc":0,"wiki":{},"dictionary":{"entry":"\u4f60\u597d","explanations":["how do you do; how are you; hello"],"related":[],"source":"wenquxing","prons":{},"type":""}}

		// 接下去要将其反序列化
		// 同样要写一个结构体，结构体的结构和返回的数据结构是一样的
		// 将相应的内容进行封装
		var dictResponse DictResponse
		json.Unmarshal(bodyText, &dictResponse)
		if err != nil {
			log.Fatal(err)
		}

		// 防御式编程
		if resp.StatusCode != 200 {
			log.Fatal("错误码：", resp.StatusCode, "body:", string(bodyText))
		}
		// 详细打印
		//fmt.Printf("%#v\n", dictResponse)
		// 打印关心内容
		// 1. 打印音标
		fmt.Println("UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
		// 2. 打印解释
		for _, item := range dictResponse.Dictionary.Explanations {
			fmt.Println(item)
		}
	}
}

func inputChinese(reader *bufio.Reader) string {
	// 读取一行的输入
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func setHeader(req *http.Request) {
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "00c1c6ca4422476f2ee9dd76573cc7df")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Chromium";v="110", "Not A(Brand";v="24", "Microsoft Edge";v="110"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.57")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")

}

/*

这段代码实现了一个HTTP POST请求，用于将一个中文字符串翻译成英文。

首先，代码创建了一个HTTP客户端实例（client），然后创建了一个包含待翻译中文字符串的字符串读取器（data）。

接着，代码使用http.NewRequest()函数创建一个HTTP请求实例（req），该实例使用POST方法发送请求，请求的URL为https://api.interpreter.caiyunai.com/v1/dict，请求的主体内容为待翻译的中文字符串。

然后，代码设置了一系列HTTP请求头，包括请求头中的数据类型、设备ID、用户代理等信息。

接下来，代码使用客户端实例的Do()方法发送HTTP请求，并接收响应（resp）。

最后，代码将响应主体读取到内存中，并输出到控制台上。在读取响应主体后，代码使用defer关键字设置了一个回收函数，以确保在程序退出前关闭响应主体的流。

*/

/*
	修改点1：封装请求体的内容，需要是手动输入的，以json的方式传递。
	修改点2：封装响应体的内容，golang并不是脚本语言，是一个强类型语言，所以也应该创建一个结构体与之对应，并将响应的内容封装在结构体中
	修改点3：输入的内容改为控制台输入
*/
