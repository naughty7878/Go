package main

import "fmt"

func main() {
	// 函数赋值给变量
	a := getSum
	fmt.Printf("a的类型：%T, getSum类型：%T\n", a, getSum)

	result := a(20, 30)
	fmt.Println("result:", result)

	// 函数作为入参
	r1 := myFun(a, 11, 22)
	fmt.Println("r1:", r1)

	// 函数类型取别名
	r2 := myFun2(a, 33, 44)
	fmt.Println("r2:", r2)

	// 支持对函数的返回值命名
	sum, sub := getSumAndSub(5, 4)
	fmt.Println("sum:", sum, "sub:", sub)

	// 支持可变参数
	total := totalSum(1, 2, 3, 4, 5)
	fmt.Println("total:", total)

}

// 支持可变参数
func totalSum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// 支持对函数的返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 函数类型取别名
type myFunType func(int, int) int

func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}
