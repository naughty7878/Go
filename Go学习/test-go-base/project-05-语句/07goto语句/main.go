package main

import "fmt"

func main() {

	// goto的使用
	fmt.Println("ok1")
	fmt.Println("ok2")
	goto label1
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
	fmt.Println("ok6")
label1:
	fmt.Println("ok7")
	fmt.Println("ok8")
	fmt.Println("ok9")
}
