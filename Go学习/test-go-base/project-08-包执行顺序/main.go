package main

import (
	"fmt"
	"myapp2/util"
)

var str string = getStr()

func getStr() string {
	fmt.Println("main.go的getStr()......")
	return "hello"
}

func init() {
	fmt.Println("main.go的init()......")
}

func main() {

	fmt.Println("main.go的main()......")
	fmt.Println("main.go的str = ", str)
	fmt.Println("main.go的util.Num = ", util.Num)

	var n1 float64 = 4.2
	var n2 float64 = 3.2
	var operator byte = '+'
	result := util.Cal(n1, n2, operator)
	fmt.Println("result = ", result)

}
