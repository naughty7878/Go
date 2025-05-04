package main

import "fmt"

func main() {

	// continue 跳出最近的一个for循环
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 3 && i == 6 {
				fmt.Print("\t\t\t")
				continue
			}
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 设置一个标签
continuei:
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 3 && i == 6 {
				fmt.Print("\t\t\t")
				fmt.Println()
				// 跳过，继续执行标签breaki的循环
				continue continuei
			}
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
