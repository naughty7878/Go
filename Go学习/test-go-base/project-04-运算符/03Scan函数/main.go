package main

import "fmt"

func main() {

	// 要求：可以从控制台接收用户信息：{姓名，年龄，薪水，是否通过考试}

	// 方式1：fmt.Scanln
	var name string
	var age int
	var sal float32
	var isPass bool

	fmt.Println("请输入姓名：")
	// 读整行
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试：")
	fmt.Scanln(&isPass)

	fmt.Printf("姓名：%v\n年龄：%v\n薪水：%v\n是否通过考试：%v\n", name, age, sal, isPass)

	// 方法二 fmt.Scanf
	fmt.Println("请输入姓名、年龄、薪水、是否通过考试（使用空格隔开）：")
	// 格式化读取
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名：%v\n年龄：%v\n薪水：%v\n是否通过考试：%v\n", name, age, sal, isPass)

}
