package main

import "fmt"

func main() {
	// ==========================================
	// 1. if — 和 Java 几乎一样，但不需要括号
	// ==========================================
	x := 10
	// Java: if (x > 5) { ... }
	// Go:   不要括号！加了不报错但不符合规范
	if x > 5 {
		fmt.Println("x > 5")
	}

	// Go 特有：if 可以带初始化语句（非常常用！）
	// 分号前是初始化，分号后是条件。y 的作用域仅限于 if-else 块内
	if y := x * 2; y > 15 {
		fmt.Println("y > 15, y =", y)
	} else {
		fmt.Println("y <= 15, y =", y)
	}
	// 这里访问不到 y，它已经超出作用域了

	// ==========================================
	// 2. for — Go 只有 for，没有 while！
	// ==========================================

	// 标准 for（和 Java 一样，但不要括号）
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 当 while 用
	// Java: while (x > 0) { ... }
	// Go:
	//n := 5
	for n := 5; n > 0; {
		fmt.Print(n, " ")
		n--
	}
	fmt.Println()

	// 无限循环
	// Java: while (true) { ... }
	// Go:
	count := 0
	for {
		if count >= 3 {
			break
		}
		count++
	}
	fmt.Println("count:", count)

	// 遍历集合（range）— 类似 Java 的 for-each
	// Java: for (String s : list) { ... }
	// Go:
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("  fruits[%d] = %s\n", i, fruit)
	}

	// 只要值，不要索引
	for _, fruit := range fruits {
		fmt.Println("  fruit:", fruit)
	}

	// ==========================================
	// 3. switch — 比 Java 的 switch 好用很多
	// ==========================================

	day := "Monday"
	// Go 的 switch 不需要 break！每个 case 自动 break
	// Java 需要手动写 break，忘了就 fall through
	switch day {
	case "Monday":
		fmt.Println("周一，上班")
	case "Saturday", "Sunday": // 可以匹配多个值
		fmt.Println("周末，休息")
	default:
		fmt.Println("工作日")
	}

	// 无条件 switch（替代 if-else if 链）

	var scoreText string
	score := 85
	switch {
	case score >= 90:
		scoreText = "优秀"
	case score >= 80:
		scoreText = "良好"
	case score >= 60:
		scoreText = "及格"
	default:
		scoreText = "不及格"
	}
	fmt.Println(scoreText)
}
