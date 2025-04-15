package main

// fmt包中提供格式化，输出，输入的函数
import "fmt"

func main() {
	// 演示转义字符
	fmt.Println("a\tb\tc")
	fmt.Println("a\nb\nc")
	fmt.Println("a\\b\\c")
	fmt.Println("a\"b\"c")
	// \r表示从当前行最前面重新输出
	fmt.Println("aaaa\rbb\rc")
}
