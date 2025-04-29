package main

import "fmt"

func main() {

	// / 除法
	fmt.Println(10 / 4)
	fmt.Println(10 / 4.0)

	// % 取模
	fmt.Println(10 % 3)

	// ++
	var i int = 10
	i++
	fmt.Println("i = ", i)
	i--
	fmt.Println("i = ", i)
	// i++ 是语句，不是表达式
	// i++（只能单独一行）
	//b := i++
	// ++i ，没有这种写法
	// i++++ 不允许多次自增, 会报错
}
