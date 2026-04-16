// ==========================================
// Go 的包 vs Java 的包
// ==========================================
//
// Java 的包：
//   - 包名和目录路径一致：com.example.service → com/example/service/
//   - 一个文件只能属于一个包
//   - import com.example.service.UserService;
//
// Go 的包：
//   - 一个目录 = 一个包（目录名通常就是包名）
//   - 同一目录下所有 .go 文件属于同一个包
//   - import "example.com/golang-web/internal/services"
//
// ==========================================
// 关键区别
// ==========================================
//
// 1. Java 按类组织，Go 按包组织
//    Java: 一个文件一个 public class
//    Go:   一个包可以有多个文件，它们共享所有内容（相当于同一个 class 拆多个文件写）
//
// 2. 包名不需要和域名对应
//    Java: package com.example.myapp.service;
//    Go:   package services（简洁，就是目录名）
//
// 3. import 路径 vs 使用时的名字
//    import "example.com/golang-web/internal/services"
//           ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//           这是 import 路径（模块路径 + 包目录）
//
//    使用时用包名：services.AuthService
//                  ^^^^^^^^
//                  这是包名（package 声明的名字）
//
// ==========================================
// 对照你的项目结构理解
// ==========================================
//
// golang-web/
// ├── go.mod                    ← 定义模块名 example.com/golang-web
// ├── main.go                   ← package main（程序入口）
// └── internal/
//     ├── domain/
//     │   └── user.go           ← package domain
//     ├── model/
//     │   └── user.go           ← package model
//     ├── repo/
//     │   └── user_repo.go      ← package repo
//     ├── services/
//     │   ├── auth_service.go   ← package services   ┐
//     │   ├── user_service.go   ← package services   ┘ 同目录 = 同一个包
//     │   └── impl/
//     │       ├── auth_service_impl.go  ← package impl
//     │       └── user_service_impl.go  ← package impl
//     └── http/
//         ├── handler/          ← package handler
//         ├── middleware/       ← package middleware
//         └── router/           ← package router
//
// Java 程序员注意：
// - internal/ 目录是 Go 的特殊约定，包外不能 import internal 下的包
//   类似 Java 的 module-info.java 限制导出（但更简单）
// - 同一个包内的文件互相访问不需要 import（它们就是一个整体）

package main

import (
	"fmt"

	// 导入标准库
	"strings"

	// 导入自己项目的包
	// 注意：import 的是路径，使用时用包名
	"example.com/golang-web/internal/domain"
)

func main() {
	// 使用标准库
	fmt.Println(strings.ToUpper("hello"))

	// 使用自己项目的包
	u := domain.User{
		// Name 大写 → 公开，包外可以访问 ✅
		// 如果 domain.User 有小写字段，这里访问不了
	}
	_ = u

	// ==========================================
	// import 别名（Java 没有的能力）
	// ==========================================
	// 如果两个包同名，可以起别名：
	// import (
	//     pkgErrors "example.com/golang-web/internal/pkg/errors"
	//     "errors"  // 标准库的 errors
	// )
	// 你项目里就这样用了：pkgErrors 是别名，避免和标准库 errors 冲突

	fmt.Println("包管理演示完成")
}
