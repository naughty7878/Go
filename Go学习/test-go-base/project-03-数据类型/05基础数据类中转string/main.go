package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num1 int = 90
	var num2 float64 = 23.345
	var b bool = true
	var c byte = 'h'
	var str string

	// 方法一：转换成string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type: %T, value : %v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type: %T, value : %v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type: %T, value : %v\n", str, str)

	str = fmt.Sprintf("%c", c)
	fmt.Printf("str type: %T, value : %v\n", str, str)

	// 方法二：使用 strconv 函数
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type: %T, value : %v\n", str, str)

	str = strconv.Itoa(num1)
	fmt.Printf("str type: %T, value : %v\n", str, str)

	str = strconv.FormatFloat(num2, 'f', -1, 32)
	fmt.Printf("str type: %T, value : %v\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str type: %T, value : %v\n", str, str)

	str = string(c)
	fmt.Printf("str type: %T, value : %v\n", str, str)

}
