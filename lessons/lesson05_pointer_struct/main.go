package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	// ==========================================
	// 1. struct 是值类型！这是和 Java 最大的区别
	// ==========================================
	//
	// Java：User u1 = new User(); User u2 = u1;
	//       u2 和 u1 指向同一个对象，改 u2 就是改 u1
	//
	// Go：  u1 := User{}; u2 := u1
	//       u2 是 u1 的完整副本，改 u2 不影响 u1

	u1 := User{Name: "Alice", Age: 25}
	u2 := u1        // 拷贝！u2 是一个全新的 User
	u2.Name = "Bob"

	fmt.Println("u1:", u1.Name) // Alice（没被改）
	fmt.Println("u2:", u2.Name) // Bob

	// ==========================================
	// 2. 想要 Java 那种"共享同一个对象"的效果 → 用指针
	// ==========================================
	u3 := &User{Name: "Alice", Age: 25} // u3 是 *User（指针）
	u4 := u3                             // u4 也指向同一个 User
	u4.Name = "Bob"

	fmt.Println("u3:", u3.Name) // Bob（被改了！因为 u3 和 u4 指向同一个对象）
	fmt.Println("u4:", u4.Name) // Bob

	// ==========================================
	// 3. 函数传参：值 vs 指针
	// ==========================================
	user := User{Name: "Alice", Age: 25}

	// 传值：函数改不了原始 user
	tryModify(user)
	fmt.Println("tryModify 后:", user.Name) // Alice

	// 传指针：函数能改原始 user
	realModify(&user)
	fmt.Println("realModify 后:", user.Name) // Modified!

	// ==========================================
	// 4. Go 的语法糖：指针访问字段不需要 *
	// ==========================================
	p := &User{Name: "Alice", Age: 25}

	// 严格写法：(*p).Name
	// Go 允许简写：p.Name（编译器自动解引用）
	fmt.Println(p.Name)     // 等价于 (*p).Name
	fmt.Println((*p).Name)  // 一样的效果，但没人这么写
}

func tryModify(u User) {
	u.Name = "Modified!" // 改的是副本
}

func realModify(u *User) {
	u.Name = "Modified!" // 通过指针改原始对象
	// 这里 u.Name 其实是 (*u).Name 的简写
}
