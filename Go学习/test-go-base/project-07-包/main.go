package main

import (
	"fmt"
	myutil "myapp/util"
)

func main() {

	var n1 float64 = 4.3
	var n2 float64 = 3.2
	var operator byte = '+'
	result := myutil.Cal(n1, n2, operator)
	fmt.Println("result = ", result)

}
