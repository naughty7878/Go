package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("mySlice类型 = %T, 长度 = %v, 容量 = %v, 值 = %v\n", mySlice, len(mySlice), cap(mySlice), mySlice)

	//fmt.Printf("mySlice[%d]地址：%p\n", 1, &mySlice[1])
	//mySlice[1] = 100
	//fmt.Printf("mySlice[%d]地址：%p\n", 1, &mySlice[1])

	// append函数追加元素
	mySlice = append(mySlice, 1000)
	fmt.Println("mySlice = ", mySlice)
	//fmt.Printf("mySlice[%d]地址：%p\n", 1, &mySlice[1])

	// 追加切片
	mySlice = append(mySlice, []int{1, 2, 2, 3, 9, 11}...)
	fmt.Println("mySlice = ", mySlice)
	//fmt.Printf("mySlice[%d]地址：%p\n", 1, &mySlice[1])

	// copy函数拷贝
	var mySlice2 []int
	copy(mySlice2, mySlice)
	fmt.Println("mySlice2 = ", mySlice2)

	var mySlice3 []int = make([]int, 10)
	copy(mySlice3, mySlice)
	fmt.Println("mySlice3 = ", mySlice3)

}
