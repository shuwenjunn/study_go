package main

import (
	"fmt"
	"os"
)

/*
学生管理系统  函数版
能够查看所有的学生  能实习增删改查
*/

type student struct {
	id   int64
	name string
}

var (
	allStudent map[int64]*student //值为student 的指针
)

func queryStudent() {
	for i, v := range allStudent {
		fmt.Printf("学号：%d，姓名：%s\n", i, v.name)
	}

}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func addStudent() {
	//1 创建一个student
	//2 追加到  allStudent Map 中
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的学号")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名")
	fmt.Scanln(&name)
	newStu := newStudent(id, name)
	allStudent[id] = newStu
	fmt.Println(allStudent)
}

func deleteStudent() {

}
func main() {

	allStudent = make(map[int64]*student, 50) //初始化开辟内存空间
	//打印菜单
	for {
		fmt.Println("欢迎光临学生管理系统")
		fmt.Println(
			`
	1.查看所有学生
	2.新增学生
	3.删除学生
	4.退出
`)
		fmt.Print("请输入你要干啥：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)
		//等待用户选择
		//执行对应的函数
		switch choice {
		case 1:
			queryStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) // 退出程序
		default:
			fmt.Println("无效输入")
		}
	}
}
