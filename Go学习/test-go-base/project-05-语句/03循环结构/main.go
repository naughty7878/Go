package main

import "fmt"

func main() {

	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//j := 0
	//for j < 10 {
	//	fmt.Println(j)
	//	j++
	//}

	k := 0
	for {
		fmt.Println(k)
		if k >= 10 {
			break
		}
		k++
	}

	var str string = "hello 北京"
	for index, elemment := range str {
		fmt.Printf("下标：%d, 值：%c\n", index, elemment)
	}
}
