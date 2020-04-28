package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//数据库信息
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/goday10"
	db, err := sql.Open("mysql", dsn) //不会校验数据库账户和密码是否正确，值校验数据源格式 ip 登
	if err != nil {                   //dsn 格式不正确会报错
		panic(err)
		return
	}
	err = db.Ping() // 尝试建立链接
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("链接数据库成功")
}
