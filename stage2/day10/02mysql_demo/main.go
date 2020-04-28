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

func queryOne(id int) {
	var u1 user
	sqlStr := `select id,name,age from user where id=?;`
	//执行
	//从连接池拿一个链接去数据库差单条记录
	//必须对rowObj 进行调用scan 方法，因为scan方法会释放链接
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	fmt.Printf("u1 %#v\n", u1)
}

func queryMore(n int) {
	sqlStr := `select id,name,age from user where id>?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		panic(err)
	}
	//特别注意
	defer rows.Close()
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("u1 %#v\n", u1)
	}
}

func inster() {
	sqlStr := `insert into user(name,age) values("王二小",88);`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		panic(err)
	}
	// 如果是插入数据的操作，能够拿到插入数据的id值
	id, err := ret.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("id:", id)
}

func updateRow(age int, id int) {
	sqlStr := `update user set age=? where id>?`
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		panic(err)
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		panic(err)
	}
	fmt.Println("更新了行数", n)

}

func deleteRow(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		panic(err)
	}
	n, err := ret.RowsAffected() //受影响的行数
	if err != nil {
		panic(err)
	}
	fmt.Println("受影响的行数", n)
}

//预处理方式插入多条数据
func pInsert() {
	sqlStr := `insert into user(name,age) values(?,?);`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
		return
	}
	//后续只需要拿到stmt去执行操作
	defer stmt.Close()
	m := map[string]int{
		"刘德华": 30,
		"周润发": 50,
		"徐闻县": 44,
		"二狗子": 33,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("链接数据库成功")
	//queryOne(1)
	//queryMore(2)
	//inster()
	//updateRow(111, 1)
	//deleteRow(5)
	pInsert()
}
