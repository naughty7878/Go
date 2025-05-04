package main

import "fmt"

func main() {

	// break 跳出最近的一个for循环
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
			if j == 3 && i == 6 {
				break
			}
		}
		fmt.Println()
	}

	// 设置一个标签
labeli:
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
			if j == 3 && i == 6 {
				// 跳出标签breaki的循环
				break labeli
			}
		}
		fmt.Println()
	}
}
