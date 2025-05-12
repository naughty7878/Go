package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串长度，len()
	// 字符和数字占1个字节，汉字占3个字节
	str := "hello中"
	fmt.Println("len(str):", len(str))

	str2 := "hello北京"
	// 字符串转成rune切片
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\t", r[i])
	}
	fmt.Println()

	// 字符串转int
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Printf("n 类型：%T, n 值：%v\n", n, n)
	}

	// int转字符串
	str3 := strconv.Itoa(123)
	fmt.Printf("str3 类型：%T, str3 值：%v\n", str3, str3)

	// 字符串转[]byte
	var bytes = []byte("hello")
	fmt.Printf("bytes 类型：%T, bytes 值：%v\n", bytes, bytes)

	// []byte转字符串
	str4 := string([]byte{97, 98, 99})
	fmt.Printf("str4 类型：%T, str4 值：%v\n", str4, str4)

	// 10进制转2进制
	str5 := strconv.FormatInt(123, 2)
	fmt.Printf("str5 类型：%T, str5 值：%v\n", str5, str5)

	// 是否包含字串
	fmt.Println(strings.Contains("hello world", "lo"))

	// 包含几个字串
	fmt.Println(strings.Count("hello worlld", "ll"))

	// 判断是否相等
	fmt.Println("\"abc\" == \"a\" + \"bc\":", "abc" == "a"+"bc")

	// 判断是否相等，不区分大小写
	fmt.Println("strings.EqualFold(\"hello\", \"heLLo\"):", strings.EqualFold("hello", "heLLo"))

	// 判断字串的第一次出现下标
	fmt.Println("strings.Index(\"hello\", \"ll\"):", strings.Index("hello", "ll"))

	// 判断字串的最后一次出现下标
	fmt.Println("strings.LastIndex(\"helloll\", \"ll\"):", strings.LastIndex("helloll", "ll"))

	// 替换
	fmt.Println("strings.Replace(\"helloll\", \"ll\", \"aa\", 1):", strings.Replace("helloll", "ll", "aa", 1))

	// 替换
	fmt.Println("strings.Split(\"a,b,c\", \",\"):", strings.Split("a,b,c", ","))

	// 转大写
	fmt.Println("strings.ToUpper(\"abc\"):", strings.ToUpper("abc"))

	// 转小写
	fmt.Println("strings.ToLower(\"ABC\"):", strings.ToLower("ABC"))

	// 去两边空白
	fmt.Printf("=====%s====\n", strings.TrimSpace(" \t\n abc \t\n "))

	// 去两边指定字符
	fmt.Printf("=====%s====\n", strings.Trim("qwabcwq", "qw"))

	// 判断前缀
	fmt.Println("strings.HasPrefix(\"qwabcwq\", \"qw\"):", strings.HasPrefix("qwabcwq", "qw"))

	// 判断后缀
	fmt.Println("strings.HasSuffix(\"qwabcwq\", \"qw\"):", strings.HasSuffix("qwabcwq", "wq"))

}
