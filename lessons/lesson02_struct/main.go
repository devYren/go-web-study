package main

import "fmt"

// ==========================================
// 1. 定义 struct — 对标 Java 的 class
// ==========================================

// Java:
// public class User {
//     private String name;
//     private int age;
//     public User(String name, int age) { ... }
// }

// Go: 没有 class，用 struct
type User struct {
	Name string
	Age  int
}

// Go 没有构造函数！惯用做法是写一个 NewXxx 工厂函数
// 对标 Java 的 new User("Alice", 25)
// 指针类型：*:表示返回是类型对象的地址，而不是对象的本身
// &表示取对象地址
func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

// ==========================================
// 2. 方法（Method）— 给 struct 添加行为
// ==========================================
// Java: public String greet() { return "Hi, " + this.name; }
// Go:   把"接收者"写在 func 和方法名之间

// (u User) 就是接收者，告诉编译器这个方法属于 User
// 值接收者：拿到的是副本，改了外面看不到（不等于 Java 的 this）
func (u User) Greet() string {
	return fmt.Sprintf("Hi, I'm %s, %d years old", u.Name, u.Age)
}

// 指针接收者：拿到的是原对象的地址，可以修改原始数据（这才等价于 Java 的 this）
func (u *User) SetAge(age int) {
	u.Age = age // 修改的是原始对象，不是副本
}

// ==========================================
// 3. 值接收者 vs 指针接收者（重点！）
// ==========================================
// Java 里所有对象都是引用传递，不存在这个问题
// Go 里你必须自己选择：
//
// 值接收者 (u User)：
//   - 方法拿到的是副本，修改不影响原对象
//   - 适合只读方法
//
// 指针接收者 (u *User)：
//   - 方法拿到的是原对象的指针，可以修改
//   - 适合需要修改状态的方法
//   - 实际开发中大部分方法用指针接收者

func main() {

	// ==========================================
	// 创建 struct 的几种方式
	// ==========================================

	// 方式 1：字面量（类似 Java 的 new + setter）
	u1 := User{Name: "Alice", Age: 25}
	fmt.Println(u1)

	// 方式 2：工厂函数（最推荐，对标 Java 构造函数）
	u2 := NewUser("Bob", 30)
	fmt.Println(*u2) // u2 是指针，*u2 取出值

	// 方式 3：零值创建（所有字段都是零值）
	var u3 User // Name="", Age=0
	fmt.Println("零值 User:", u3)

	// ==========================================
	// 调用方法
	// ==========================================
	fmt.Println(u1.Greet())

	// 指针接收者的方法
	fmt.Println("修改前:", u2.Age)
	u2.SetAge(31)
	fmt.Println("修改后:", u2.Age)

	// ==========================================
	// 值接收者的证明：修改不影响原对象
	// ==========================================
	u4 := User{Name: "Charlie", Age: 20}
	u4.tryChange()                      // 值接收者，改不了
	fmt.Println("tryChange 后:", u4.Age) // 还是 20
}

func (u User) tryChange() {
	u.Age = 999 // 改的是副本，外面看不到
}
