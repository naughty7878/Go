package main

import "fmt"

func main() {
	// & 运算符（取地址）
	var num int = 42
	ptr := &num      // ptr 是一个指向 num 的指针，类型为 *int
	fmt.Println(ptr) // 输出类似: 0xc00000a0b8（内存地址）

	// * 运算符（解引用）
	fmt.Println(*ptr) // 输出: 42（解引用获取值）

	*ptr = 100       // 通过指针修改值
	fmt.Println(num) // 输出: 100
}
