package main

import "fmt"

func main() {
	// golang的变量使用方式

	// 方法一：golang 一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1 = ", n1, ", n2 = ", n2, ", n3 = ", n3)

	// 方法二：golang 一次性声明多个变量
	var num, name = 100, "tom"
	fmt.Println("num = ", num, ", name = ", name)

	// 方法三：golang 一次性声明多个变量
	v1, v2 := 200, "jack"
	fmt.Println("v1 = ", v1, ", v2 = ", v2)
}
