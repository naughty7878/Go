package main

import (
	"fmt"
	"time"
)

func main() {

	//test()

	safeCall()

	for {
		fmt.Println("=============")
		time.Sleep(1 * time.Second)
	}

}

func test() {
	// 使用 defer + recover 来捕获和处理异常
	defer func() {
		// recover()内置函数，可以捕获异常
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res = ", res)
}

func safeCall() {
	defer func() {
		// recover()内置函数，可以捕获异常
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// panic() 抛出运行时异常
	panic("something bad happened")
}
