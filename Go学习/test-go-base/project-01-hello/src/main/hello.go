// 把这个test.go 归属到 main
package main

// 引入一个包 fmt
import "fmt"

// 主方法
// func 是一个关键字，表示函数
// main 是函数名
func main() {
	// 打印
	fmt.Println("hello world")

	// fmt.Printnl("hello world")
}

// 编译命令：go build hello.go

// 执行：./hello

// 或者： go run hello.go

// 注意事项
/*
1) Go 源文件以 "go" 为扩展名。
2) Go应用程序的执行入口是main()函数。这个是和其它编程语言(比如java/c)
3) Go语言严格区分大小写。
4) Go方法由一条条语句构成，每个语句后不需要分号(Go语言会在每行后自动加分号)，这也体现出 Golang 的简洁性。
5) Go编译器是一行行进行编译的,因此我们一行就写一条语句，不能把多条语句写在同一个，否则报错
6) go语言定义的变量或者import的包如果没有使用到,代码不能编译通过
7) 大括号都是成对出现的,缺一不可
*/
