package main

import "fmt"

func main() {

	// 编写一个程序，输入年龄，
	// 如果年龄大于等于18，则输出“年龄大于等于18岁，已成年”
	fmt.Println("请输入年龄：")
	var age int
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("年龄大于等于18岁，已成年")
	} else {
		fmt.Println("未成年")
	}

	// 支持if语句中直接定义变量，作用域只在if结构中
	if age1 := 17; age1 >= 18 {
		fmt.Println("年龄大于等于18岁，已成年")
	} else {
		fmt.Println("未成年")
	}
}
