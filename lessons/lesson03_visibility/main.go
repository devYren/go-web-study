package main

import "fmt"

// ==========================================
// Go 的可见性规则：大写 = public，小写 = private
// 就这一条规则，没有 public/private/protected 关键字
// ==========================================

// Java:
// public class User { ... }          ← public 关键字
// private String name;               ← private 关键字
// protected int age;                 ← protected 关键字

// Go: 首字母大写 = 包外可见（public），首字母小写 = 包内可见（private）
type User struct {
	Name string // 大写：包外可见（public）
	age  int    // 小写：仅本包可见（private）
}

// 大写方法：包外可以调用（public）
func (u *User) GetAge() int {
	return u.age
}

// 小写方法：仅本包内可以调用（private）
func (u *User) validate() bool {
	return u.Name != "" && u.age > 0
}

// 大写函数：包外可以调用（public 工厂函数）
func NewUser(name string, age int) *User {
	return &User{Name: name, age: age}
}

// 小写函数：仅本包内可以调用
func helper() string {
	return "我是内部工具函数"
}

func main() {
	u := NewUser("Alice", 25)

	// 同一个包内，大写小写都能访问
	fmt.Println(u.Name)       // ✅ 大写，公开
	fmt.Println(u.age)        // ✅ 小写，但同一个包内可以访问
	fmt.Println(u.GetAge())   // ✅ 大写方法
	fmt.Println(u.validate()) // ✅ 小写方法，同包内可以
	fmt.Println(helper())     // ✅ 小写函数，同包内可以

	// 如果从另一个包访问：
	// u.Name      ✅ 可以
	// u.age       ❌ 编译错误！小写字段包外不可见
	// u.GetAge()  ✅ 可以
	// u.validate() ❌ 编译错误！
}
