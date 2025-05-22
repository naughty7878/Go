package main

import "fmt"

func main() {

	myArr := [5]int{76, 25, 55, 23, 74}

	// 冒泡排序
	BubbleSort(&myArr)
	// 二分查找
	BinarySearch(&myArr, 74)
}

// 冒泡排序
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前：arr=", *arr)
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j+1] = (*arr)[j+1] + (*arr)[j]
				(*arr)[j] = (*arr)[j+1] - (*arr)[j]
				(*arr)[j+1] = (*arr)[j+1] - (*arr)[j]
			}
		}
	}
	fmt.Println("排序后：arr=", *arr)
}

// 二分查找
func BinarySearch(arr *[5]int, num int) {
	fmt.Printf("查找源数组：%v, 查找的值：%v\n", *arr, num)
	var left, right int = 0, len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if (*arr)[mid] > num {
			right = mid - 1
		} else if (*arr)[mid] < num {
			left = mid + 1
		} else {
			fmt.Printf("查找的值：%v, 下标：%v\n", num, mid)
			return
		}
	}
	fmt.Println("未找到对应的zhi")
}
