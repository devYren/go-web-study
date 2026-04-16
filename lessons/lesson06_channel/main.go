package main

import "fmt"

// ==========================================
// Channel（通道）— goroutine 之间传递数据
// ==========================================
//
// Java 中 goroutine 通信要用 BlockingQueue、Future、CompletableFuture 等
// Go 只需要一个内置类型：channel
//
// Go 的并发哲学：
//   "不要通过共享内存来通信，而要通过通信来共享内存"
//   翻译成 Java 话：别用 synchronized 锁共享变量，用 channel 传消息

func main() {
	// ==========================================
	// 1. 基本用法
	// ==========================================
	// 创建一个传递 string 的 channel
	// 对标 Java: BlockingQueue<String> ch = new LinkedBlockingQueue<>(0);
	ch := make(chan string)

	// goroutine 往 channel 发送数据
	go func() {
		ch <- "hello" // 发送（会阻塞，直到有人接收）
	}()

	// 主 goroutine 从 channel 接收数据
	msg := <-ch // 接收（会阻塞，直到有人发送）
	fmt.Println("收到:", msg)

	// ==========================================
	// 2. Channel 是阻塞的（这是关键！）
	// ==========================================
	// ch <- value  发送方阻塞，等接收方准备好
	// value := <-ch  接收方阻塞，等发送方发数据
	//
	// 不需要锁！channel 自带同步机制

	// ==========================================
	// 3. 带缓冲的 Channel
	// ==========================================
	// 无缓冲：发送和接收必须同时准备好（像打电话，双方必须同时在线）
	// 有缓冲：可以先发送几条，接收方稍后取（像发短信，先存着）
	bufferedCh := make(chan int, 3) // 缓冲区大小 3
	bufferedCh <- 1                 // 不阻塞（缓冲区没满）
	bufferedCh <- 2                 // 不阻塞
	bufferedCh <- 3                 // 不阻塞
	// bufferedCh <- 4              // 这里会阻塞！缓冲区满了

	fmt.Println(<-bufferedCh) // 1
	fmt.Println(<-bufferedCh) // 2
	fmt.Println(<-bufferedCh) // 3

	// ==========================================
	// 4. 实际场景：用 channel 收集 goroutine 的结果
	// ==========================================
	fmt.Println("\n=== 收集结果 ===")
	results := make(chan string, 3)

	apis := []string{"/users", "/orders", "/products"}
	for _, api := range apis {
		go func(url string) {
			results <- fmt.Sprintf("response from %s", url)
		}(api)
	}

	// 接收 3 个结果
	for i := 0; i < len(apis); i++ {
		fmt.Println(<-results)
	}

	// ==========================================
	// 5. 关闭 channel + range 遍历
	// ==========================================
	fmt.Println("\n=== range 遍历 channel ===")
	numCh := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numCh <- i
		}
		close(numCh) // 发完了，关闭 channel
	}()

	// range 会一直读取，直到 channel 被关闭
	for n := range numCh {
		fmt.Print(n, " ")
	}
	fmt.Println()
}
