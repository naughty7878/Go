package main

import "fmt"

// 演示golang中的指针
func main() {

	var num int = 9
	fmt.Printf("num 地址：%v\n", &num)

	var ptr *int = &num
	fmt.Printf("ptr 的值：%v\n", ptr)

	// 修改指针值对应的地址的值
	// 会导致num值的变化
	*ptr = 20
	fmt.Printf("num 的值：%v\n", num)
}
