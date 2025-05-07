package main

import "fmt"

func main() {
	// 从终端输入打印对应乘法表
	fmt.Println("请输入需要打印几阶乘法表：")
	var n int
	fmt.Scanln(&n)
	fmt.Printf("现在开始打印%d阶乘法表\n", n)
	printMulti(n)
}

func printMulti(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
