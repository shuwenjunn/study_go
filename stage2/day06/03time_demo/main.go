package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) // 时间戳纳秒

	//time.Unix() 根据时间戳返回时间

	ret := time.Unix(1581060680, 0)
	fmt.Println(ret)
	fmt.Println(ret.Day())
	fmt.Println(ret.Month())
	fmt.Println(ret.Year())

	//时间间隔
	fmt.Println(time.Second)

	fmt.Println(now.Add(10 * time.Minute))

	//定时器
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t)
	//}

	//时间格式化
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM"))

	//字符串的时间转化成时间戳

	timeObj, err := time.Parse("2006-01-02", "1998-07-27")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(timeObj.Unix())
}
