package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("now 类型=%T，now 值=%v\n", now, now)

	// 获取年月日时分秒
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	ms := time.Now().UnixNano() / 1000000
	unixTimestamp := time.Now().Unix()
	millis := time.Now().UnixMilli()
	fmt.Printf("year 类型=%T，值=%v\n", year, year)
	fmt.Printf("month 类型=%T，值=%v\n", month, month)
	fmt.Printf("month 类型=%T，值=%v\n", month, int(month))
	fmt.Printf("day 类型=%T，值=%v\n", day, day)
	fmt.Printf("hour 类型=%T，值=%v\n", hour, hour)
	fmt.Printf("minute 类型=%T，值=%v\n", minute, minute)
	fmt.Printf("second 类型=%T，值=%v\n", second, second)
	fmt.Printf("ms 类型=%T，值=%v\n", ms, ms)
	fmt.Printf("unixTimestamp 类型=%T，值=%v\n", unixTimestamp, unixTimestamp)
	fmt.Printf("millis 类型=%T，值=%v\n", millis, millis)

	// 格式化一
	fmt.Printf("时间：%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// 格式化二
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format(time.DateTime))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format(time.DateOnly))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format(time.TimeOnly))

	for i := 0; i < 100; i++ {
		fmt.Println(time.Now().Format(time.DateTime))
		// 休眠
		time.Sleep(1 * time.Second)
	}
}
