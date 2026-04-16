package main

import "fmt"

// ==========================================
// panic / recover — Go 的最后手段
// ==========================================
//
// Go 日常用 error 返回值处理错误，但也有类似 Java throw 的机制：
//
//   panic  → 类似 Java 的 throw new RuntimeException（程序崩溃）
//   recover → 类似 Java 的 catch（捕获 panic，阻止崩溃）
//
// 重要原则：
//   - 99% 的错误用 error 返回值处理
//   - panic 只在"程序不可能继续运行"时用（如启动时配置缺失）
//   - recover 几乎只在框架层用（如 Gin 的 Recovery 中间件）
//   - 业务代码里基本不要用 panic/recover

func mustLoadConfig() {
	// 假设配置文件不存在，程序没法跑
	panic("config file not found, cannot start")
	// panic 后面的代码不会执行（类似 throw 后面的代码）
}

// 用 Java 类比整个 safeCall：
//
// void safeCall() {
// try {
// mustLoadConfig();                    // 2. 这里 throw
// System.out.println("这行不会执行");
// } catch (Exception r) {                  // 3. 捕获
// System.out.println("捕获到: " + r);
// }
// // finally 也是不管有没有异常都执行，和 defer 一样
// }
//
// defer 就是 Go 版的 try-finally，recover 就是 Go 版的 catch，只是写法不同。Go 把它们拆成了两个独立的概念。
func safeCall() {
	// defer + recover 可以捕获 panic，防止程序崩溃,对标 Java 的 try-catch

	//defer = 延迟执行，意思是"这行代码等函数结束时在执行"
	defer func() {
		//recover()的捕获范围是当前函数
		if r := recover(); r != nil {
			fmt.Println("捕获到 panic:", r)
		}
	}() // 这个 () 是调用！表示"定义完立刻调用"

	mustLoadConfig() // 这里会 panic,所以下面那行不会执行
	fmt.Println("这行不会执行")
}

func main() {
	// ==========================================
	// 你项目里的实际用法 — Recovery 中间件
	// ==========================================
	// internal/http/middleware/recover.go 就是做这件事：
	//
	// func Recovery() gin.HandlerFunc {
	//     return func(c *gin.Context) {
	//         defer func() {
	//             if rec := recover(); rec != nil {
	//                 // 把 panic 转成 JSON 错误响应，而不是让服务器崩溃
	//             }
	//         }()
	//         c.Next()
	//     }
	// }
	//
	// 这就是为什么 Web 服务不会因为一个请求 panic 而整个挂掉

	safeCall()
	fmt.Println("程序继续运行（没有崩溃）")

	// 如果不 recover，panic 会导致程序直接退出
	// 取消下面的注释试试：
	mustLoadConfig() // 程序直接崩溃退出
}
