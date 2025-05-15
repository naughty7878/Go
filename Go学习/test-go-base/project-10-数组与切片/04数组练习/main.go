package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 打印 A-Z
	var myChars [26]byte
	for i := 0; i < len(myChars); i++ {
		myChars[i] = 'A' + byte(i)
	}
	for i := 0; i < len(myChars); i++ {
		fmt.Printf("myChars[%v] = %c\n", i, myChars[i])
	}

	// 求数组最大值
	var intArr = [...]int{1, 2, 3, 4, 5}
	for i := 1; i < len(intArr); i++ {
		if intArr[0] < intArr[i] {
			intArr[0] += intArr[i]
			intArr[i] = intArr[0] - intArr[i]
			intArr[0] = intArr[0] - intArr[i]
		}
	}
	fmt.Printf("intArr = %v\n", intArr)

	// 求平均值 与 和
	var intArr2 = [...]int{1, 2, 3, 4, 6}
	sum := 0
	for _, value := range intArr2 {
		sum += value
	}
	fmt.Printf("sum = %v\n", sum)
	fmt.Printf("sum / len(intArr2) = %v\n", float64(sum)/float64(len(intArr2)))

	// 随机生成5个数，并反转
	var intArr3 [5]int
	// 1. 设置随机种子（确保每次运行结果不同）
	//rand.Seed(time.Now().UnixNano())
	// 1. 创建自定义的随机数生成器（替代 rand.Seed）
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < len(intArr3); i++ {
		// // 生成随机 int（范围：0 ~ 100）
		intArr3[i] = r.Intn(100)
	}
	fmt.Printf("intArr3 = %v\n", intArr3)
	// 交换
	for i := 0; i < len(intArr3)/2; i++ {
		intArr3[i] += intArr3[len(intArr3)-1-i]
		intArr3[len(intArr3)-1-i] = intArr3[i] - intArr3[len(intArr3)-1-i]
		intArr3[i] = intArr3[i] - intArr3[len(intArr3)-1-i]
	}
	fmt.Printf("intArr3 = %v\n", intArr3)
}
