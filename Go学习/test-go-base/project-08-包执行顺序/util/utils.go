package util

import "fmt"

var Num int = getNum()

func getNum() int {
	fmt.Println("utils.go的getNum()......")
	return 100
}

func init() {
	fmt.Println("utils.go的init()......")
}

func Cal(n1 float64, n2 float64, operator byte) float64 {

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
