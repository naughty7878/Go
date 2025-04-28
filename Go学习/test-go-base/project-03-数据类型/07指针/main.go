package main

import "fmt"

// 演示golang中的指针
func main() {

	// 基础数据类型在内存中的布局
	var i int = 10
	// 获取 i 的地址，使用 & 取地址符
	fmt.Println("i 地址是：", &i)

	// 指针的值是变量的内存地址
	var ptr *int = &i
	fmt.Println("ptr 值是：", ptr)
	//fmt.Println("ptr 地址是：", &ptr)
	fmt.Println("ptr 指向的值是：", *ptr)
}
