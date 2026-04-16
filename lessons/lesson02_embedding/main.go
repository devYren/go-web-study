package main

import "fmt"

// ==========================================
// Go 没有继承！用组合（Embedding）替代
// ==========================================

// Java 的做法：
// class Animal { String name; void eat() { ... } }
// class Dog extends Animal { void bark() { ... } }

// Go 的做法：把一个 struct 嵌入另一个 struct
type Base struct {
	Name string
}

func (b Base) Eat() {
	fmt.Println(b.Name, "is eating")
}

// Dog "嵌入" Base，获得 Base 的所有字段和方法
type Dog struct {
	Base        // 嵌入（不写字段名，只写类型）
	Breed string
}

// Dog 可以有自己的方法
func (d Dog) Bark() {
	fmt.Println(d.Name, "says: 汪汪！") // 直接访问 Base 的 Name
}

// ==========================================
// 对照项目里的 gorm.Model
// ==========================================
// 你项目里的 UserModel 就用了嵌入：
//
// type UserModel struct {
//     gorm.Model           ← 嵌入！自动获得 ID、CreatedAt、UpdatedAt、DeletedAt
//     Username     string
//     Email        string
//     PasswordHash string
// }
//
// 不需要继承 BaseEntity，直接嵌入就获得了所有公共字段

func main() {
	d := Dog{
		Base:  Base{Name: "旺财"},
		Breed: "柴犬",
	}

	d.Eat()  // 调用的是 Base 的方法，不需要 d.Base.Eat()
	d.Bark() // Dog 自己的方法

	// 也可以直接访问嵌入 struct 的字段
	fmt.Println("名字:", d.Name)   // 等价于 d.Base.Name
	fmt.Println("品种:", d.Breed)
}
