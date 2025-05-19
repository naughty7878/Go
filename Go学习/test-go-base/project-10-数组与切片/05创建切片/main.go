package main

import "fmt"

func main() {

	// 方法1
	var arr = [5]int{1, 2, 3, 4, 5}
	mySlice := arr[2:5]
	fmt.Printf("mySlice类型 = %T, 长度 = %v, 容量 = %v, 值 = %v\n", mySlice, len(mySlice), cap(mySlice), mySlice)

	// 方法二
	mySlice2 := make([]int, 5)
	fmt.Printf("mySlice2类型 = %T, 长度 = %v, 容量 = %v, 值 = %v\n", mySlice2, len(mySlice2), cap(mySlice2), mySlice2)

	mySlice3 := make([]int, 5, 10)
	fmt.Printf("mySlice3类型 = %T, 长度 = %v, 容量 = %v, 值 = %v\n", mySlice3, len(mySlice3), cap(mySlice3), mySlice3)

	// 方法三
	mySlice4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("mySlice4类型 = %T, 长度 = %v, 容量 = %v, 值 = %v\n", mySlice4, len(mySlice4), cap(mySlice4), mySlice4)

}
