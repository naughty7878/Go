package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("mySlice类型 = %T, 长度 = %v, 容量 = %v, 值 = %v\n", mySlice, len(mySlice), cap(mySlice), mySlice)

	// 方式一：for循环遍历
	for i := 0; i < len(mySlice); i++ {
		fmt.Printf("mySlice[%d] = %d\n", i, mySlice[i])
	}

	// 方式二：for-rang遍历
	for index, value := range mySlice {
		fmt.Printf("index = %d, value = %v\n", index, value)
	}
}
