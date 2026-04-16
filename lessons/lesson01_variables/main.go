package main

import "fmt"

func main() {
	// ==========================================
	// 1. 变量声明：Go vs Java
	// ==========================================

	// Java: String name = "Alice";
	// Go 写法一：完整声明（类似 Java）
	var name string = "Alice"

	// Go 写法二：类型推断（类似 Java 的 var，但 Go 从诞生就有）
	var age = 25

	// Go 写法三：短声明 := （最常用！只能在函数内部用）
	email := "alice@example.com"

	fmt.Println(name, age, email)

	// ==========================================
	// 2. 类型对照表
	// ==========================================
	// Java int       → Go int（但 Go 还有 int8, int16, int32, int64）
	// Java long      → Go int64
	// Java double    → Go float64
	// Java boolean   → Go bool
	// Java String    → Go string
	// Java 没有      → Go uint（无符号整数）

	var score float64 = 99.5
	var passed bool = true
	fmt.Println(score, passed)

	// ==========================================
	// 3. 零值（Zero Value）— 这点和 Java 很像
	// ==========================================
	// Java 的 int 默认 0，boolean 默认 false，引用类型默认 null
	// Go 也有默认值，但叫"零值"：
	var x int     // 0
	var s string  // ""（空字符串，不是 null！Go 没有 null）
	var b bool    // false
	fmt.Printf("零值: int=%d, string=%q, bool=%t\n", x, s, b)

	// ==========================================
	// 4. 常量
	// ==========================================
	// Java: static final String VERSION = "1.0";
	// Go:
	const version = "1.0"
	fmt.Println("version:", version)

	// ==========================================
	// 5. 多变量声明
	// ==========================================
	// Java 没有这种写法，Go 可以一行声明多个
	a, b2, c := 1, "hello", true
	fmt.Println(a, b2, c)

	// 交换两个变量（Java 需要临时变量，Go 不用！）
	x1, x2 := 10, 20
	x1, x2 = x2, x1
	fmt.Println("交换后:", x1, x2) // 20, 10
}
