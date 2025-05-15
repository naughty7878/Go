package main

import (
	"fmt"
)

func main() {

	// 使用数组
	//var score [5]float64
	//
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("请输入第%d个元素的值\n", i+1)
	//	fmt.Scan(&score[i])
	//}
	//
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("arr[%d] = %v\n", i, score[i])
	//}

	// 4种初始化方式
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{1, 2, 3}
	// ... 是规定写法不能取消
	var arr3 = [...]int{1, 2, 3}
	// 指定元素下标
	var arr4 = [3]string{1: "tom", 0: "ab", 2: "qq"}
	fmt.Println("arr1 = ", arr1)
	fmt.Println("arr2 = ", arr2)
	fmt.Println("arr3 = ", arr3)
	fmt.Println("arr4 = ", arr4)

	for index, value := range arr1 {
		fmt.Printf("元素下标：%d, 元素值：%v\n", index, value)
	}

	// 默认零值
	fmt.Println("=============默认零值===========")
	var arr5 [3]int
	for index, value := range arr5 {
		fmt.Printf("元素下标：%d, 元素值：%v\n", index, value)
	}
}
