package main

import (
	"errors"
	"fmt"
)

// ==========================================
// 1. 基本函数
// ==========================================
// Java: public int add(int a, int b) { return a + b; }
// Go:   类型写在变量名后面（这是最需要适应的地方）
func add(a int, b int) int {
	return a + b
}

// 参数类型相同时可以简写
func multiply(a, b int) int {
	return a * b
}

// ==========================================
// 2. 多返回值 — Go 最重要的特性之一！
// ==========================================
// Java 只能返回一个值，想返回多个要包装成对象
// Go 天然支持多返回值，这是 Go 错误处理的基石

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil // nil 类似 Java 的 null，但只用于特定类型
}

// ==========================================
// 3. 命名返回值（Named Return）
// ==========================================
func swap(a, b int) (x int, y int) {
	x = b
	y = a
	return // 裸 return，自动返回命名的变量
}

// ==========================================
// 4. 函数是一等公民
// ==========================================
// Java 需要 Function<T,R> 或 lambda，Go 的函数天然就是值
func apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {

	// 基本调用
	fmt.Println("3 + 4 =", add(3, 4))
	fmt.Println("3 * 4 =", multiply(3, 4))

	// 多返回值 + 错误处理（Go 的核心模式！）
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	// 故意触发错误
	// 用 _ 忽略不需要的返回值（Java 没有这个概念）
	// _ 是"空白标识符"，表示"我不关心这个值"
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("错误:", err) // 输出: 错误: division by zero
	}

	// 命名返回值
	x, y := swap(1, 2)
	fmt.Printf("swap(1,2) = %d, %d\n", x, y)

	// 函数作为参数
	fmt.Println("apply add:", apply(3, 4, add))
	fmt.Println("apply multiply:", apply(3, 4, multiply))

	// 匿名函数（类似 Java 的 lambda）
	fmt.Println("apply lambda:", apply(3, 4, func(a, b int) int {
		return a - b
	}))

}
