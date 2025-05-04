package main

import "fmt"

var age int = getAge()

func main() {
	fmt.Println("执行main()......")
	fmt.Println("age = ", age)
}

// init函数，通常用于完成初始化工作
func init() {
	fmt.Println("执行init()......")
}

func getAge() int {
	fmt.Println("执行getAge()......")
	return 90
}
