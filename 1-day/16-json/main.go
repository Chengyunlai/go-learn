package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user := User{Name: "Chengyunlai", Age: 24}
	fmt.Println(user) //{Chengyunlai 24}
	// 转json
	res, err := json.Marshal(user)
	fmt.Println(res)         //[123 34 78 97 109 101 34 58 34 67 104 101 110 103 121 117 110 108 97 105 34 44 34 65 103 101 34 58 50 52 125]
	fmt.Println(string(res)) //{"Name":"Chengyunlai","Age":24}
	fmt.Println(err)
	// 反序列化
	user2 := User{}
	json.Unmarshal(res, &user2)
	fmt.Println(user2) //{Chengyunlai 24}

}

// 1 要求结构体的变量首字母为大写
type User struct {
	Name string `json:"name"` // 当转json时指定key的名称
	Age  int    `json:"age"`
}
