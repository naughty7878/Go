package main

import "fmt"

// 定义全局变量
var n1 = 100
var name1 = "tom"

// 方式二
var (
	n2    = 200
	name2 = "jack"
)

func main() {

	// 使用全局变量
	fmt.Println("n1 = ", n1, "name1 = ", name1)

	fmt.Println("n2 = ", n2, "name2 = ", name2)
}
