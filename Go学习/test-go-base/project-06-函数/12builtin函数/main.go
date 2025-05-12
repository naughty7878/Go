package main

import (
	"fmt"
)

func main() {
	num1 := 100
	fmt.Printf("num1类型: %T, num1值: %v\n", num1, num1)

	// new 给值类型分配内存
	num2 := new(int)
	fmt.Printf("num2类型: %T, num2值: %v\n", num2, num2)
	fmt.Printf("num2指针对应的值：%v\n", *num2)

	// make 给引用类型分配内存
}
