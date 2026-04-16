package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ==========================================
// 1. Java vs Go 的错误处理哲学
// ==========================================
//
// Java：异常是"非正常流程"，用 try-catch 捕获
//   try {
//       int n = Integer.parseInt(str);
//   } catch (NumberFormatException e) {
//       // 处理错误
//   }
//
// Go：错误就是普通的返回值，用 if err != nil 检查
//   n, err := strconv.Atoi(str)
//   if err != nil {
//       // 处理错误
//   }
//
// Go 没有 try-catch！错误必须显式处理，不能忽略。

func main() {
	// ==========================================
	// 2. 标准模式：函数返回 (结果, error)
	// ==========================================

	// strconv.Atoi = Java 的 Integer.parseInt
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("解析失败:", err)
	} else {
		fmt.Println("解析成功:", n)
	}

	// 故意传错误的值
	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("解析失败:", err) // 解析失败: strconv.Atoi: parsing "abc": invalid syntax
	}

	// ==========================================
	// 3. 创建错误的几种方式
	// ==========================================

	// 方式 1：errors.New（最简单，类似 new Exception("msg")）
	err1 := errors.New("something went wrong")
	fmt.Println(err1)

	// 方式 2：fmt.Errorf（可以格式化，最常用）
	name := "test.txt"
	err2 := fmt.Errorf("failed to open file: %s", name)
	fmt.Println(err2)

	// ==========================================
	// 4. 实际开发中的错误处理模式
	// ==========================================
	content, err := readFile("config.yaml")
	if err != nil {
		fmt.Println("读取失败:", err)
		// Java 里你可能 throw 或 catch
		// Go 里你 return err 向上传播，或者就地处理
	} else {
		fmt.Println("内容:", content)
	}

	// ==========================================
	// 5. 错误链（Error Wrapping）— 类似 Java 的 cause
	// ==========================================
	err = processConfig()
	if err != nil {
		fmt.Println("最终错误:", err)
		// 输出: 最终错误: process config: read config: file not found
		// 层层包装，保留完整链路，排查问题很方便
	}
}

// 模拟一个会失败的函数
func readFile(name string) (string, error) {
	// 模拟文件不存在
	return "", fmt.Errorf("file not found: %s", name)
}

// ==========================================
// 错误包装：用 %w 把原始错误包进去
// ==========================================
// Java 的做法：throw new ServiceException("msg", cause);
// Go 的做法：fmt.Errorf("msg: %w", err)

func readConfig() error {
	return fmt.Errorf("read config: %w", errors.New("file not found"))
	//                          ^^
	//                          %w 是关键！它把原始错误包进去
	//                          后面可以用 errors.Is / errors.As 解包
}

func processConfig() error {
	err := readConfig()
	if err != nil {
		// 再包一层，保留原始错误
		return fmt.Errorf("process config: %w", err)
	}
	return nil
}
