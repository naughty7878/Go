package main

import "fmt"

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}

func sum(n1 int, n2 int) int {
	// 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立栈（defer栈）
	// 当函数执行完成后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=", n1)
	// 在defer将语句放入栈时，也会将相关的值拷贝同时入栈，即n2的值不管
	defer fmt.Println("ok2 n2=", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}
