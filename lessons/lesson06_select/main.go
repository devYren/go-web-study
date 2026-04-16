package main

import (
	"fmt"
	"time"
)

// ==========================================
// select — 同时等待多个 channel
// 对标 Java 的 CompletableFuture.anyOf() 或 NIO Selector
// ==========================================

func main() {
	// ==========================================
	// 1. 基本 select
	// ==========================================

	//chan = Channel类型， chan string 表示存放string类型的通道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// goroutine 1：100ms 后往 ch1 发数据
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自 ch1"
	}()

	// goroutine 2：50ms 后往 ch2 发数据
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "来自 ch2"
	}()

	// select 同时监听多个 channel，谁先来就处理谁
	// 类似 switch，但用于 channel
	select {
	case msg := <-ch1: // ch1 有数据了就走这里
		fmt.Println("收到:", msg)
	case msg := <-ch2: // ch2 有数据了就走这里
		fmt.Println("收到:", msg) // ch2 更快（50ms），所以先到
	}

	// ==========================================
	// 2. 最常用场景：超时控制
	// ==========================================
	fmt.Println("\n=== 超时控制 ===")
	result := make(chan string)

	//  ch <- "hello"      // channel 在左边：发送（把数据塞进去）
	//  msg := <-ch        // channel 在右边：接收（把数据取出来）
	go func() {
		time.Sleep(2 * time.Second) // 模拟慢操作
		result <- "操作完成"
	}()

	select {
	case msg := <-result:
		fmt.Println(msg)
	case <-time.After(500 * time.Millisecond): // time.After返回的就是一个存Time类型的Channel
		fmt.Println("超时了！") // 500ms 内没收到结果，走这里
	}

	// ==========================================
	// 3. 对照项目中的实际用法
	// ==========================================
	// 你项目 main.go 里的 graceful shutdown 就用了类似的模式：
	//
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// sig := <-quit  // 阻塞，直到收到关闭信号
	//
	// 这里 quit 是一个 channel，程序一直阻塞等待
	// 当你按 Ctrl+C 时，操作系统发送 SIGINT 信号到这个 channel
	// <-quit 收到信号，程序开始优雅关闭
}
