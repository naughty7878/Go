package main

import "fmt"

// 全局匿名函数
var Fun1 = func(n1 int, n2 int) int {
	return n1 + n2
}

func main() {

	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1 = ", res1)

	a := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Printf("a 类型：%T\n", a)
	fmt.Println("a(1, 2) 调用： ", a(1, 2))

	// 调用全局匿名函数
	fmt.Println("Fun1(1, 2) 调用： ", Fun1(1, 2))

}
