package main

import "fmt"

func main() {

	// 二维数组
	// 3个班 每个版5个同学 成绩
	var scores [3][5]float64

	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生成绩\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	// 平均成绩
	totalSum := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += float64(scores[i][j])
		}
		fmt.Printf("第%d班的总分%v，平均分%v\n", i+1, sum, sum/float64(len(scores[i])))
		totalSum += sum
	}
	fmt.Printf("总分%v，平均分%v\n", totalSum, totalSum/float64(len(scores)))
}
