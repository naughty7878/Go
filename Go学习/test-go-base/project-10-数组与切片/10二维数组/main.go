package main

import "fmt"

func main() {

	// 二维数组
	var arr [4][6]int
	fmt.Printf("arr类型：%T，值：%v\n", arr, arr)

	// 赋值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	// 遍历
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}

	// 地址
	fmt.Printf("arr地址：%p\n", &arr)
	fmt.Printf("arr[0]地址：%p\n", &arr[0])
	fmt.Printf("arr[1]地址：%p\n", &arr[1])
	fmt.Printf("arr[2]地址：%p\n", &arr[2])

	fmt.Printf("arr[0][0]地址：%p\n", &arr[0][0])
	fmt.Printf("arr[1][0]地址：%p\n", &arr[1][0])

	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {3, 4, 5}}
	fmt.Printf("arr2类型：%T，值：%v\n", arr, arr)
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Print(arr2[i][j], "\t")
		}
		fmt.Println()
	}

	// 遍历 for-rang
	for _, item := range arr2 {
		//fmt.Println(index, item)
		for _, value := range item {
			fmt.Print(value, "\t")
		}
		fmt.Println()
	}
}
