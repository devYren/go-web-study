// ==========================================
// go.mod 对照 Java 的 pom.xml
// ==========================================
//
// Java (pom.xml):
//   <groupId>com.example</groupId>
//   <artifactId>my-app</artifactId>
//   <version>1.0</version>
//   <dependencies>
//       <dependency>
//           <groupId>org.springframework</groupId>
//           <artifactId>spring-web</artifactId>
//           <version>5.3.0</version>
//       </dependency>
//   </dependencies>
//
// Go (go.mod):
//   module example.com/golang-web     ← 模块名（类似 groupId + artifactId）
//   go 1.26.1                         ← Go 版本
//   require (
//       github.com/gin-gonic/gin v1.12.0   ← 依赖（类似 dependency）
//       gorm.io/gorm v1.25.12
//   )
//
// ==========================================
// 常用命令对照
// ==========================================
//
// Java (Maven)              Go
// ─────────────────────     ────────────────────
// mvn install               go mod tidy（整理依赖）
// mvn dependency:resolve    go mod download（下载依赖）
// mvn clean package         go build（编译）
// mvn exec:java             go run .（运行）
// mvn test                  go test ./...（测试）
//
// ==========================================
// 关键区别
// ==========================================
//
// 1. Go 没有中央仓库（没有 Maven Central / Nexus）
//    依赖直接从 Git 仓库下载：github.com/gin-gonic/gin
//    国内可以用代理：GOPROXY=https://goproxy.cn
//
// 2. Go 的依赖存在全局缓存里（$GOPATH/pkg/mod）
//    不像 Maven 每个项目有自己的 target/，Go 所有项目共享缓存
//
// 3. go.sum 类似 Maven 的 dependency lock
//    记录了每个依赖的哈希值，确保构建可重复
//
// 4. 添加依赖超级简单：
//    直接在代码里 import，然后运行 go mod tidy，自动下载加到 go.mod
//    不需要手动编辑 go.mod（不像 Java 要手动改 pom.xml）

package main

import "fmt"

func main() {
	fmt.Println("go.mod 讲解完成，请阅读本文件顶部的注释")
}
