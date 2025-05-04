package main

import "fmt"

func main() {

	var n1 float64 = 4.3
	var n2 float64 = 3.2
	var operator byte = '+'
	result := cal(n1, n2, operator)
	fmt.Println("result = ", result)

}

func cal(n1 float64, n2 float64, operator byte) float64 {

	var result float64
	switch operator {
	case '+':
		result = n1 + n2
	case '-':
		result = n1 - n2
	case '*':
		result = n1 * n2
	case '/':
		result = n1 / n2
	default:
		fmt.Println("操作符号错误...")
	}
	return result
}
