package main

import (
	"fmt"
	"time"
)

//全局变量，init函数和main函数执行顺序依次是全局变量--init函数--main函数
var age = test()

func test() int {

	fmt.Println("我被赋值给全局变量了")
	return 29
}

func init() {
	fmt.Println("我是inti函数，一般执行初始化的操作")
}

//短写法必须用在方法内  不是全局变量
// name := "hello"
var name string = "name"

func main() {

	//map的定义
	var aa map[int]string
	aa = make(map[int]string)
	fmt.Printf("aa=%T\n", aa)

	bb := make(map[int]string)
	fmt.Printf("bb=%T\n", bb)

	dd := map[int]string{}
	fmt.Printf("cc=%T\n", dd)

	println(name)
	timeDemo2()
	c()
	a()
	var number1 int = 99
	var number2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string

	//%d按十进制
	str = fmt.Sprintf("%d", number1)
	fmt.Printf("str type %T str=%q\n", str, str)
	//%f按浮点数
	str = fmt.Sprintf("%f", number2)
	fmt.Printf("str type %T str=%q\n", str, str)
	//%t布尔占位符
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)
	//%c字符占位符
	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str=%q\n", str, str)

	var a int = 1 >> 2
	var b1 int = -1 >> 2
	var c int = 1 << 2
	var d int = -1 << 2
	//a,b,c,d 结果是多少
	fmt.Println("a=", a)
	fmt.Println("b=", b1)
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	//字符串中包含汉字，该遍历形式可以输出汉字
	var bj string = "hello world, 北京！"
	for i, strbj := range bj {
		fmt.Printf("i=%d strbj=%c\n", i, strbj)
	}
	//字符串中包含汉字，该遍历形式不可以输出汉字，会出现乱码
	for i := 0; i < len(bj); i++ {
		fmt.Printf("i=%d strbj=%c\n", i, bj[i])
	}
	//字符串中包含汉字，该遍历形式可以输出汉字（将字符串转换成[]rune切片）
	bj2 := []rune(bj)
	for i := 0; i < len(bj2); i++ {
		fmt.Printf("i=%d strbj=%c\n", i, bj2[i])

	}

}

type circle struct {
	radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func add(int) int {

	return 10
}

func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return
	// return a, b
}

func namedRetValus() (int, int) {

	return 1, 2
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
