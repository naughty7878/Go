package main

import "fmt"

func main() {

	// 切片使用
	var mySlice []int
	fmt.Printf("mySlice类型 = %T, mySlice = %v\n", mySlice, mySlice)
	fmt.Printf("len(mySlice) = %v\n", len(mySlice))
	fmt.Printf("cap(mySlice) = %v\n", cap(mySlice))

	// 添加元素
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 5)

	// 删除元素 下标为2的元素
	mySlice = append(mySlice[:2], mySlice[2+1:]...)

	fmt.Printf("mySlice类型 = %T, mySlice = %v\n", mySlice, mySlice)
	fmt.Printf("len(mySlice) = %v\n", len(mySlice))
	fmt.Printf("cap(mySlice) = %v\n", cap(mySlice))

	for i := 0; i < len(mySlice); i++ {
		fmt.Println("mySlice[i] = ", mySlice[i])
	}

	mySlice2 := mySlice
	mySlice2[1] = 99
	fmt.Printf("mySlice = %v\n", mySlice)
	fmt.Printf("mySlice2 = %v\n", mySlice2)

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr地址 = %p\n", &arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %p\n", i, &arr[i])
	}
	mySlice3 := arr[1:]
	fmt.Printf("mySlice3地址 = %p\n", &mySlice3)
	fmt.Printf("mySlice3[0]地址 = %p\n", &mySlice3[0])

}
