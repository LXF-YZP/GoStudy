package main

import "fmt"

func main() {

	//类型断言
	// var x interface{}
	// var b2 float32 = 1.1
	// x = b2 //空接口可以接受任何类型
	// y := x.(float32)
	// fmt.Printf("y 的类型是%T 值是=%v\n", y, x)

	//类型断言（带有检查报错）
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //空接口可以接受任何类型
	// y, ok := x.(float32)
	// if ok {
	// 	fmt.Printf("y 的类型是%T 值是=%v\n", y, x)
	// } else {
	// 	fmt.Println("类型断言错误")
	// }
	if y, ok := x.(float32); ok {
		fmt.Printf("y 的类型是%T 值是=%v\n", y, x)
	} else {
		fmt.Println("类型断言错误")
	}

}
