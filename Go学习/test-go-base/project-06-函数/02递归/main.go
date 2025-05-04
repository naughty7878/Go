package main

import "fmt"

func main() {

	fmt.Println("第一天桃子数量", peach(1))

}

func peach(n int) int {
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}
