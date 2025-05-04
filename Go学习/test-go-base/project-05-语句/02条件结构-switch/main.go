package main

import "fmt"

func main() {

	// 编写一个程序，接收一个字符，
	// a表示星期一，b表示星期二，根据用户的输入显示响应的信息

	var key byte
	fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)
	switch key + 1 {
	case 'a':
		fmt.Println("星期一")
	case 'b' + 10:
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期天")
	default:
		fmt.Println("输入有误！")
	}

	var age int = 10
	switch {
	case age == 10:
		fmt.Println("年龄10")
	case age == 20:
		fmt.Println("年龄20")
	default:
		fmt.Println("年龄不满足")
	}

	switch age2 := age; {
	case age2 == 10:
		fmt.Println("年龄10")
		fallthrough
	case age2 == 20:
		fmt.Println("年龄20")
	default:
		fmt.Println("年龄不满足")
	}
}
