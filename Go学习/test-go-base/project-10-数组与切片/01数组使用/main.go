package main

import (
	"fmt"
	"math"
)

func main() {

	// 使用数组
	var arr [6]float64
	arr[0] = 3.0
	arr[1] = 5.0
	arr[2] = 1.0
	arr[3] = 3.4
	arr[4] = 2.0
	arr[5] = 50.0

	totalWeigh := 0.0
	for i := 0; i < len(arr); i++ {
		totalWeigh += arr[i]
	}

	avWeight := RoundToTwoDecimalPlaces(totalWeigh / float64(len(arr)))
	fmt.Printf("totalWeigh = %v, avWeight = %v\n", totalWeigh, avWeight)

	// 数组地址就是数组第一个元素地址
	fmt.Printf("arr类型 = %T\n", arr)
	fmt.Printf("arr地址 = %p\n", &arr)
	fmt.Printf("arr[0]地址 = %p\n", &arr[0])
	fmt.Printf("arr[1]地址 = %p\n", &arr[1])
}

// RoundToTwoDecimalPlaces 四舍五入保留2位小数
func RoundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}
