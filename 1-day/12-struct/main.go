package main

import "fmt"

func main() {
	// 直接初始化时指定
	user := User{name: "chengyunlai", password: "root"}
	fmt.Println(user) // {chengyunlai root}
	// 修改名称
	user.name = "Cheng"
	fmt.Println(user) // {Cheng root}

	// 定义一个变量再指定属性
	user2 := User{}
	user2.name = "GoLang"
	fmt.Println(user2) //{GoLang }

	fmt.Println(checkPassword(&user, "123"))
	fmt.Println(user.checkPassword("123"))

}

type User struct {
	name     string
	password string
}

// 指针可以减少拷贝开销，也可以修改原值
func checkPassword(u *User, pass string) bool {
	return (*u).password == pass
}

// 结构体方法
func (u *User) resetPassword(password string) {
	u.password = password
}

func (u *User) resetUserName(userName string) {
	u.name = userName
}

func (u *User) checkPassword(pass string) bool {
	return (*u).password == pass
}
