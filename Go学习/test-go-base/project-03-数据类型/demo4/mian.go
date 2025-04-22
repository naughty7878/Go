package main

import "fmt"

func main() {

	var i int32 = 100
	// i => float
	var n1 float32 = float32(i)
	fmt.Printf("n1 = %v, 类型：%T", n1, n1)
}
