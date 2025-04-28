package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str string = "true"
	var b bool
	// 在 Go 语言中，_ 是 空白标识符（Blank Identifier），
	// 它的作用是忽略某个返回值的占位符
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type: %T, value : %v\n", b, b)

	var str2 string = "12345"
	var n1 int
	n1, _ = strconv.Atoi(str2)
	fmt.Printf("n1 type: %T, value : %v\n", n1, n1)

	var n2 int64
	n2, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n2 type: %T, value : %v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type: %T, value : %v\n", f1, f1)

	// 转换失败
	var str4 string = "hello"
	var n4 int
	var err error
	n4, err = strconv.Atoi(str4)
	fmt.Printf("n4 type: %T, value : %v\n", n4, n4)
	fmt.Printf("err type: %T, value : %v\n", err, err)

}
