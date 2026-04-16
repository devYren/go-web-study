package main

import (
	"fmt"
	"time"
)

// ==========================================
// Goroutine vs Java Thread
// ==========================================
//
// Java 创建线程：
//   new Thread(() -> {
//       System.out.println("hello from thread");
//   }).start();
//
// Go 创建 goroutine：
//   go func() {
//       fmt.Println("hello from goroutine")
//   }()
//
// 就一个关键字：go
//
// 核心区别：
//   Java Thread：操作系统级线程，每个占用约 1MB 内存，创建几千个就很吃力
//   Go Goroutine：Go 运行时管理的轻量级协程，每个只占约 2KB，轻松创建几十万个

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: 第 %d 次\n", name, i+1)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// ==========================================
	// 1. 普通调用（顺序执行）
	// ==========================================
	fmt.Println("=== 顺序执行 ===")
	sayHello("A")
	sayHello("B")
	// A 跑完才跑 B，总共需要约 600ms

	// ==========================================
	// 2. 用 go 关键字并发执行
	// ==========================================
	fmt.Println("\n=== 并发执行 ===")
	go sayHello("A") // 启动 goroutine，不等它完成
	go sayHello("B") // 再启动一个 goroutine

	// 重要！main 函数不会等 goroutine 完成
	// 如果 main 退出了，所有 goroutine 立刻被杀掉
	// 这里先用 time.Sleep 等一下（后面会讲正确的等待方式）
	time.Sleep(500 * time.Millisecond)

	// ==========================================
	// 3. go + 匿名函数
	// ==========================================
	fmt.Println("\n=== 匿名函数 ===")
	go func() {
		fmt.Printf("我是匿名 goroutines")
	}()

	// 带参数的匿名 goroutine
	msg := "hello"
	go func(s string) {
		fmt.Println("收到:", s)
	}(msg) // 注意这里传参，不要直接引用外部变量（闭包陷阱）

	time.Sleep(100 * time.Millisecond)
}
