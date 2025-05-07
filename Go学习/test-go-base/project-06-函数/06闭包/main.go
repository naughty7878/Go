package main

import (
	"fmt"
	"strings"
)

func main() {
	// 使用累加器
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	// 使用makeSuffix
	m := makeSuffix(".jpg")
	fmt.Println(m("winter"))
	fmt.Println(m("bird.jpg"))
}

// 累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

// 判断后最
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}
