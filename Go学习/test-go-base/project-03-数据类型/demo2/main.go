package main

import "fmt"

func main() {
	// 引号-字符串
	var str = "hello"
	fmt.Println("str = ", str)

	// 反引号-字符串
	var str2 = `var str string = "hello"
	fmt.Println("str = ", str)`
	fmt.Println("str2 = ", str2)

	// 加号-拼接
	var str3 = "hello" + "world"
	fmt.Println("str3 = ", str3)

	// 加号-换行
	var str4 = "hello" + "world" + "world" + "world" + "world" +
		"world" + "world" + "world"
	fmt.Println("str4 = ", str4)
}
