package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	file, err := os.Open("/Users/luxiaofeng/Develop/go/src/GoStudy/src/familyaccount/main/main.go")
	if err != nil {
		fmt.Println("读取文件有问题")
	}
	fmt.Printf("file=%v\n", file)
	err = file.Close()
	if err != nil {
		fmt.Println("关闭文件有问题")
	}
	//iostat -x 3
	//sar -n DEV 3
	var tt T
	fmt.Println(tt.A)
	fmt.Println(tt.B)
	fmt.Println(tt.C)
	fmt.Println(tt.D)
	data, err := json.Marshal(tt)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))

	//分割字符串
	str := "223,344,"
	fmt.Println(strings.Contains(str, ","))
	s := strings.Split(str, ",")
	for j := 0; j < len(s); j++ {
		fmt.Println(s[j])
	}

	//判断map值是否为空
	myMap := make(map[string]string, 3)
	var myKey = "123"
	myMap[myKey] = "520"
	value, ok := myMap[myKey]
	if ok {
		//存在
		fmt.Println(value)
	} else {
		//不存在
		fmt.Println("不存在")
	}

	//打印程序执行时间
	start := time.Now()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

	//1、使用关键字var，声明变量类型并赋值
	var v1 int = 100
	//2、使用关键字var，直接对变量赋值，go可以自动推导出变量类型
	var v2 = 10
	//3、直接使用“：=”对变量赋值，不使用var，两者同时使用会语法冲突，推荐使用
	v3 := 10
	fmt.Println(v1, v2, v3)
	fmt.Println(v4, v5)
	//交换两个值不需要第三个变量

	v1, v2 = v2, v1

	fmt.Println(v1, v2)
	_, _, nickName := GetName() //使用匿名变量丢弃部分返回值
	fmt.Println(nickName)
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Wednesday)
	fmt.Println(numberOfDays)

	//1、浮点型分为float32(类似C中的float)，float64(类似C中的double)
	var f1 float32
	f1 = 12 //不加小数点，被推导为整型
	fmt.Printf("f1=%v, f1=%T\n", f1, f1)
	f2 := 12.0       //加小数点，被推导为float64
	f1 = float32(f2) //需要执行强制转换
	//2、浮点数的比较
	//浮点数不是精确的表达方式，不能直接使用“==”来判断是否相等，可以借用math的包math.Fdim
	fmt.Printf("f1=%v, f1=%T\n", f1, f1)
	fmt.Printf("f1=%v, f1=%T\n", f2, f2)

	//go语言中字符类型

	var myArray [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("myArray=%v, myArray=%T\n", myArray, myArray)
	slice01 := myArray[:]
	fmt.Printf("slice01=%v, slice01=%T\n", slice01, slice01)

	slice02 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("slice02=%v, slice02=%T\n", slice02, slice02)

	slice03 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("slice03=%v, slice03=%T\n", slice03, slice03)

	var mySlice []int = myArray[:]

	fmt.Printf("mySlice=%v, mySlice=%T\n", mySlice, mySlice)

	var mySlice01 []int = []int{1, 2, 3}
	fmt.Printf("mySlice01=%v, mySlice01=%T\n", mySlice01, mySlice01)

	newSlice := append(mySlice01, 1, 2, 3) //直接将元素加进去，若存储空间不够会按上述方式扩容。
	fmt.Printf("newSlice=%v, newSlice=%T\n", newSlice, newSlice)
	newSlice1 := append(newSlice, mySlice01...) //将oldSlice2的元素打散后加到oldSlice1中，三个点不可省略。
	fmt.Printf("newSlice1=%v, newSlice1=%T\n", newSlice1, newSlice1)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	// copy(slice2, slice1) //只会复制slice1的前三个元素到slice2中
	copy(slice1, slice2) //只会复制slice2的三个元素到slice1中的前三个位置
	// fmt.Printf("slice1=%v, slice1=%T\n", slice1, slice1)
	fmt.Printf("slice2=%v, slice2=%T\n", slice2, slice2)

	//new 和make区别
	p := new([]int) //p == nil; with len and cap 0
	fmt.Println(p)

	v := make([]int, 10, 50) // v is initialed with len 10, cap 50
	fmt.Println(v)

	/*********Output****************
	    &[]
	    [0 0 0 0 0 0 0 0 0 0]
	*********************************/

	// (*p)[0] = 18 // panic: runtime error: index out of range
	// because p is a nil pointer, with len and cap 0
	v[1] = 18 // ok
	fmt.Println(v)

}

//var mySlice []int=myArray[first:last]
const (
	Sunday = iota //Sunday==0,以此类推
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday     //大写字母开头表示包外可见
	numberOfDays //小写字母开头表示包内私有
)

//Go中所有声明后的变量都需要调用到，当出现函数多返回值，并且部分返回值不需要使用时，可以使用匿名变量丢弃该返回值
func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Make"
}

//1、使用关键字var，声明变量类型并赋值
var v4 int

//2、使用关键字var，直接对变量赋值，go可以自动推导出变量类型
var v5 = 10

// //3、直接使用“：=”对变量赋值，不使用var，两者同时使用会语法冲突，推荐使用
// v3 := 10
// fmt.Println(v1, v2, v3)

type T struct {
	A bool
	B int    "myb"           //默认0
	C string "myc,omitempty" //默认""
	D string `bson:",omitempty" json:"jsonkey"`
}
