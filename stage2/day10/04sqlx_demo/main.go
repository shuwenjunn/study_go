package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB //db 就是数据库的连接池对象

func initDB() (err error) {
	//数据库信息
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/sql_test"
	db, err = sqlx.Connect("mysql", dsn) //不会校验数据库账户和密码是否正确，值校验数据源格式 ip 登
	if err != nil {                      //dsn 格式不正确会报错
		return
	}
	//db.SetMaxOpenConns(10) //设置数据库连接池的最大连接数
	//db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	return
}

type user struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var u user
	err := initDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("链接数据库成功")
	err = db.Get(&u, `select id,name,age from user where id=1;`)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	var userList = make([]user, 0, 10)
	err = db.Select(&userList, `select id,name,age from user`)
	fmt.Printf("userList %v\n", userList)
}
