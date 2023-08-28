package main

import (
	"database/sql"
)

func main() {
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Juejin")
	if err != nil {
		panic(err)
	}

}
