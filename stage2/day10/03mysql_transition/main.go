package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //db 就是数据库的连接池对象

func initDB() (err error) {
	//数据库信息
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn) //不会校验数据库账户和密码是否正确，值校验数据源格式 ip 登
	if err != nil {                  //dsn 格式不正确会报错
		return
	}
	err = db.Ping() // 尝试建立链接
	if err != nil {
		return
	}

	//db.SetMaxOpenConns(10) //设置数据库连接池的最大连接数
	//db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	return
}

type user struct {
	id   int
	name string
	age  int
}

func transactionDemo() {
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		panic(err)
		return
	}

	// 执行多个操作
	sqlStr1 := `update user set age=age-2 where id=1;`
	sqlStr2 := `update user set age=age+2 where id=2;`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Println("1 fail")
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Println("2 fail")
		return
	}
	// 上面两步都执行成功了，提交本次事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}
	fmt.Println("事务执行成功")
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("链接数据库成功")
	transactionDemo()
}
