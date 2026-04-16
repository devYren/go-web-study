package main

import (
	"fmt"
	"sync"
	"time"
)

// ==========================================
// 上一节用 time.Sleep 等待 goroutine，这是错误的做法
// 正确做法：sync.WaitGroup
// 对标 Java 的 CountDownLatch
// ==========================================

func main() {
	// ==========================================
	// 1. WaitGroup 基本用法
	// ==========================================
	//
	// Java (CountDownLatch):
	//   CountDownLatch latch = new CountDownLatch(3);
	//   for (int i = 0; i < 3; i++) {
	//       new Thread(() -> { doWork(); latch.countDown(); }).start();
	//   }
	//   latch.await(); // 等待所有线程完成
	//
	// Go (WaitGroup):
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 计数器 +1（"还有一个任务要做"）
		go func(id int) {
			defer wg.Done() // 计数器 -1（"我做完了"），defer 确保一定执行
			fmt.Printf("worker %d 开始\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("worker %d 完成\n", id)
		}(i)
	}

	wg.Wait() // 阻塞，直到计数器归零（所有 goroutine 都 Done 了）
	fmt.Println("所有 worker 完成")

	// ==========================================
	// 2. 实际场景：并发请求多个 API
	// ==========================================
	fmt.Println("\n=== 并发请求 ===")
	var wg2 sync.WaitGroup
	apis := []string{"/users", "/orders", "/products"}

	for _, api := range apis {
		wg2.Add(1)
		go func(url string) {
			defer wg2.Done()
			result := fetchAPI(url)
			fmt.Println(result)
		}(api)
	}

	wg2.Wait()
	fmt.Println("所有 API 请求完成")
}

func fetchAPI(url string) string {
	time.Sleep(100 * time.Millisecond) // 模拟网络延迟
	return fmt.Sprintf("response from %s", url)
}
