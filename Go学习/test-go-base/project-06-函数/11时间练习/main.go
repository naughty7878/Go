package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now().Unix()
	test()
	endt := time.Now().Unix()
	fmt.Println("执行方法时间(s):", endt-start)
}

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
