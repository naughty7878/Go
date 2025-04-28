package main

import "fmt"

func main() {

	// 值类型
	a := 10
	b := a            // b 是 a 的副本
	b = 20            // 修改 b 不影响 a
	fmt.Println(a, b) // 输出: 10 20

	// 值类型（结构体）
	type Person struct{ Name string }
	p1 := Person{"Alice"}
	p2 := p1        // p2 是 p1 的副本
	p2.Name = "Bob" // 不影响 p1
	fmt.Println(p1) // {Alice}

	// 值类型（数组）
	a1 := [3]int{1, 2, 3} // 定义一个数组
	a2 := a1              // a2 是 a1 的完整副本，不是引用！

	a2[0] = 99      // 修改 a2 不会影响 a1
	fmt.Println(a1) // 输出: [1 2 3]（原数组未变）
	fmt.Println(a2) // 输出: [99 2 3]

	// 切片（引用类型）
	s1 := []int{1, 2, 3}
	s2 := s1        // s2 和 s1 共享底层数组
	s2[0] = 99      // 修改 s2 会影响 s1
	fmt.Println(s1) // 输出: [99 2 3]

	// 映射（引用类型）
	m1 := map[string]int{"a": 1}
	m2 := m1        // m2 和 m1 共享底层数据
	m2["a"] = 100   // 修改 m2 会影响 m1
	fmt.Println(m1) // 输出: map[a:100]
}
