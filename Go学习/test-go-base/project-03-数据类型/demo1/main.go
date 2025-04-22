package main

import (
	"fmt"
	"unsafe"
)

// golang中整数类型的使用
func main() {

	i := 1
	fmt.Println("i=", i)

	// 测试一下int8的范围
	//var j int8 = -129   // 编译报错
	//fmt.Println("j=", j)

	// 查看类型
	fmt.Printf("i=%T\n", i)

	// 查看占用字节数
	fmt.Printf("i占用字节：%d\n", unsafe.Sizeof(i))

	var j byte = 1
	fmt.Printf("j的类型是： %T, j占用的字节：%d;\n", j, unsafe.Sizeof(j))

	k := 3.14
	k2 := .14
	fmt.Printf("k的类型是： %T, k占用的字节：%d;\n", k, unsafe.Sizeof(k))
	fmt.Printf("k2的类型是： %T, k2占用的字节：%d;\n", k2, unsafe.Sizeof(k2))
	fmt.Println("k2 = ", k2)

	c1 := 'a'
	c2 := "a"
	fmt.Printf("c1的类型是： %T, c1占用的字节：%d, 值是：%c, 码值是：%d;\n", c1, unsafe.Sizeof(c1), c1, c1)
	fmt.Printf("c2的类型是： %T, c2占用的字节：%d, 值是：%s;\n", c2, unsafe.Sizeof(c2), c2)
}
