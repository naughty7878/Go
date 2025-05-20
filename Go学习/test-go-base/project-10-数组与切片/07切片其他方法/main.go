package main

import "fmt"

func main() {

	// string底层是一个byte数组
	str := "hello world"

	strSlice := []byte(str[6:])
	fmt.Printf("strSlice类型：%T, 值： = %v\n", strSlice, strSlice)

	list := []byte(str)
	list[0] = 'z'
	str2 := string(list)
	fmt.Println("str2 =", str2)

	// 斐波那契数列
	fmt.Println("fbc(10) =", fbc(10))
}

func fbc(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
