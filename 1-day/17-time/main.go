package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) //2023-05-12 17:32:12.0876215 +0800 CST m=+0.005319101

	// 构造时间
	t1 := time.Date(2023, 05, 12, 17, 32, 0, 0, time.UTC)
	t2 := time.Date(2023, 05, 12, 18, 32, 0, 0, time.UTC)
	fmt.Println(t1) //2023-05-12 17:32:00 +0000 UTC

	// 方法
	fmt.Println(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute())
	fmt.Println(t1.Format("2006-01-02 15:04:05")) //2023-05-12 17:32:00（无语这个格式。）
	diff := t2.Sub(t1)
	fmt.Println(diff)                           //1h0m0s
	fmt.Println(diff.Minutes(), diff.Seconds()) // 60 3600

	t3, err := time.Parse("2006-01-02 15:04:05", "2023-05-12 17:39:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3)
	// 获取时间戳
	fmt.Println(now.Unix()) // 1683884463

}
