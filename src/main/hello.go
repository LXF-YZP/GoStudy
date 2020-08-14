package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//os.Args demo
func main() {

	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}

	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	timeDemo()

	// log.Fatalln("这是一条会触发fatal的日志。")
	// log.Panicln("这是一条会触发panic的日志。")

	// 通过 `name(args)` 来调用函数，
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	// res = plusPlus(1, 2, 3)
	// fmt.Println("1+2+3 =", res)
	a, b := vals()
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	sum(1, 2)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	//元素的类型和长度都是数组类型的一部分
	//数组默认是零值的，对于int数组来说也就是0
	var a1 [5]int
	fmt.Println("emp", a1)
	a1[4] = 100
	fmt.Println("set:", a1)
	fmt.Println("get:", a1[4])
	fmt.Println("len:", len(a1))
	b1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b1)
	for i := 0; i < len(b1); i++ {
		fmt.Print(b1[i], " ")
	}
	var two [2][3]int
	fmt.Println("two.len:", len(two))
	for i := 0; i < len(two); i++ {
		for j := 0; j < len(two[0]); j++ {
			two[i][j] = i + j
		}
	}
	fmt.Println("2d:", two)
	//slice的类型仅仅由它所包含的元素决定（不需要元素个数）
	//要创建一个长度非零的空slice，需要使用内建的方法make
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len：", len(s))

	s = append(s, "d")
	fmt.Println("ass:", s)
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	l := s[2:5]
	fmt.Println("sli:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//匿名函数可以赋值给变量
	// sum01 := func() int {
	// 	sum := 0
	// 	for i := 1; i <= 1e6; i++ {
	// 		sum += i
	// 	}
	// 	return sum
	// }

	// fmt.Println(sum01())

	//匿名函数计算1到1百万的和
	func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
		fmt.Println(sum)
	}()

	ss := intSeq()
	println(ss())
	fmt.Println(ss())
	sss := plusPlus(1)
	println(sss)
	aaa := sub(1, 2)
	println(aaa)

	bb := facts(7)
	println(bb)

	//值类型和指针类型
	println("--------------------------------------------------------------")
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// 通过 `&i` 语法来取得 `i` 的内存地址，即指向 `i` 的指针。
	fmt.Println("pointer:", &i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针也是可以被打印的。
	fmt.Println("pointer:", &i)

	//函数做为参数
	add(1, 2)
	callback(1, add)

	testo()

	array := [3]float64{7.0, 8.5, 9.1}
	x := ssum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)

	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	println(cap(slice1), "--------------")
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	//键盘输入

	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
}

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func ssum(a *[3]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}

func testo() {
	var i int
	println("-----------------------", i)
}

func add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

// 我们将通过两个函数：`zeroval` 和 `zeroptr` 来比较指针和值类型的不同。
// `zeroval` 有一个 `int` 型参数，所以使用值传递。
// `zeroval` 将从调用它的那个函数中得到一个 `ival` 形参的拷贝。
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` 有一和上面不同的 `*int` 参数，意味着它用了一个 `int` 指针。
// 函数体内的 `*iptr` 接着_解引用_这个指针，从它内存地址得到这个地址对应的当前值。
// 对一个解引用的指针赋值将会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
	*iptr = 0
}

//
func facts(a int) int {
	if a == 0 {
		return 1
	}
	return a * facts(a-1)
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func plusPlus(int) int {
	a := 1
	b := 2
	c := 3
	return a + b + c
}

//闭包测试
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 这个函数接受任意数目的 `int` 作为参数。
func sum(nums ...int) {

	total := 0
	for i, num := range nums {
		total += num
		fmt.Println(i)
	}
	fmt.Println(total)
	fmt.Println(nums)
}

// `(int, int)` 在这个函数中标志着这个函数返回 2 个 `int`。
func vals() (int, int) {
	return 3, 7
}

func timeDemo() {
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

// 这里是一个函数，接受两个 `int` 并且以 `int` 返回它们的和
func plus(a int, b int) int {

	// Go 需要明确的返回，不会自动返回最
	// 后一个表达式的值
	return a + b
}

// 当多个连续的参数为同样类型时，最多可以仅声明最后一个参数类型
// 而忽略之前相同类型参数的类型声明。
// func plusPlus(a, b, c int) int {
// 	return a + b + c
// }
