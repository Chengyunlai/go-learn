package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

/*
 1. 系统生成随机数
    包："math/rand"
    rand.Intn(x)生成随机数：[0,maxNum)之间的随机数
*/
func main() {
	maxNum := 10000
	// 要点1：使用rand.Intn(x)生成一个[0,100)之间的随机数
	// Go 1.20以前惯例添加一个随机数种子，使其更加有随机性，而现在已经无需再添加随机数种子了
	// rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum) // [0,maxNum)
	// 打印一下随机数
	//println("随机数：", secretNumber)
	reader := bufio.NewReader(os.Stdin)
	// 一共10次机会
	num := 10
	for i := 0; i < num; i++ {
		guess, err := inputNum(reader)
		if err != nil {
			fmt.Println(err)
			fmt.Println("你还有", num-i-1, "次机会,一共", num, "次机会")
			continue
		}
		fmt.Println("你猜测的数字是", guess, "你还有", num-i-1, "次机会,一共", num, "次机会")
		if guess > secretNumber {
			fmt.Println("你猜测的数字太大了")
		} else if guess < secretNumber {
			fmt.Println("你猜测的数字太小了")
		} else {
			fmt.Println("恭喜你猜对了")
			break
		}
	}
}

func inputNum(reader *bufio.Reader) (int, error) {
	fmt.Println("请猜测一个整型数字")
	// 读取一行输入
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("在读取你输入的过程中发生了一个错误，请稍后重试", err)
		return 0, err
	}
	// 去掉换行符，TrimSuffix提供返回一个不含"\r\n"的字符串。注意Linux换行符为：\n 而Windows的换行符是：\r\n
	input = strings.TrimSuffix(input, "\r\n")
	// 试图将字符串数字转换成数字
	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("非法输入，请输入一个整型数字", err)
		return 0, err
	}
	return guess, nil
}

/**
案例一：猜测随机数

内容：系统生成一个随机数[0~100)，用户可以在键盘中输入数字并猜测该数字是什么，只有10次猜测机会。
*/
