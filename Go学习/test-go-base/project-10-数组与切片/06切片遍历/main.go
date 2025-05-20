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

	// 切片的切片
	mySlice2 := mySlice[1:2]
	fmt.Printf("mySlice2类型 = %T, 长度 = %v, 容量 = %v, 值 = %v\n", mySlice2, len(mySlice2), cap(mySlice2), mySlice2)
	mySlice2[0] = 100
	fmt.Printf("mySlice类型 = %T, 长度 = %v, 容量 = %v, 值 = %v\n", mySlice, len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("mySlice2类型 = %T, 长度 = %v, 容量 = %v, 值 = %v\n", mySlice2, len(mySlice2), cap(mySlice2), mySlice2)

	// mySlice 与 mySlice2 指向的是同一个数组的空间，所有改变某个值，其他引用也会改变
}
