package main

import (
	"fmt"
	"time"
)

func main() {
	timeDemo2()
	c()
	a()
}

func a() {
	i := 0
	defer fmt.Println(0) //因为i=0，所以此时就明确告诉golang在程序退出时，执行输出0的操作
	i++
	return
}

func c() (i int) {
	defer func() {
		i++
		println(i)
	}()
	return 1
}

func timeDemo2() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
