package main

import "fmt"

// ==========================================
// 1. Java 的 interface vs Go 的 interface
// ==========================================

// Java 的做法（显式声明实现关系）：
// public interface Animal { String Speak(); }
// public class Dog implements Animal { ... }  ← 必须写 implements
// public class Cat implements Animal { ... }  ← 必须写 implements

// Go 的做法（隐式实现！不需要声明 implements）：
type Animal interface {
	Speak() string
}

// Dog 只要有 Speak() string 方法，就自动实现了 Animal 接口
// 不需要写任何 implements 关键字！
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + ": 汪汪！"
}

// Cat 也是，有 Speak() string 就自动实现了 Animal
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + ": 喵喵！"
}

// 这个函数接受任何实现了 Animal 接口的类型
// 对标 Java 的 void makeSound(Animal a) { ... }
func makeSound(a Animal) {
	fmt.Println(a.Speak())
}

// ==========================================
// 2. 为什么隐式实现是好设计？
// ==========================================
// Java 的问题：如果你要让第三方库的类实现你的接口，做不到
//   （你改不了别人的代码，加不了 implements）
//
// Go 的优势：只要方法签名匹配，自动就实现了
//   你可以为任何类型定义接口，不需要对方配合

// ==========================================
// 3. 实际项目中的用法 — 对照你的 golang-web 项目
// ==========================================
// 在你的项目里，UserRepository 就是一个接口：
//
// type UserRepository interface {
//     Create(ctx context.Context, u *model.UserModel) error
//     FindByID(ctx context.Context, id uint) (*model.UserModel, error)
//     ...
// }
//
// userRepo struct 实现了这些方法，所以它自动就是 UserRepository
// Service 层只依赖接口，不关心具体实现
// 这就是为什么可以轻松替换实现或写测试 mock

// ==========================================
// 4. 空接口 interface{} — 对标 Java 的 Object
// ==========================================
// 4. 空接口 any — 对标 Java 的 Object
// ==========================================
// Go 1.18 之前写 interface{}，之后可以写 any（是同一个东西）
//
//	 这两个完全一样
//	func printAnything(v any) { ... }
//	func printAnything(v interface{}) { ... }
func printAnything(v any) {
	fmt.Printf("类型: %T, 值: %v\n", v, v)
}

// ==========================================
// 5. 类型断言 — 对标 Java 的 instanceof + 强制转换
// ==========================================
func describeAnimal(a Animal) {
	// Java: if (a instanceof Dog) { Dog d = (Dog)a; ... }
	// Go:
	switch v := a.(type) {
	case Dog:
		fmt.Println("这是一只狗:", v.Name)
	case Cat:
		fmt.Println("这是一只猫:", v.Name)
	default:
		fmt.Println("未知动物")
	}
}

func main() {
	// 多态（和 Java 一模一样的效果）
	dog := Dog{Name: "旺财"}
	cat := Cat{Name: "咪咪"}

	makeSound(dog)
	makeSound(cat)

	// 接口类型的变量可以持有任何实现者
	var a Animal
	a = dog
	fmt.Println(a.Speak())
	a = cat
	fmt.Println(a.Speak())

	// 空接口：任何类型都满足
	printAnything(42)
	printAnything("hello")
	printAnything(dog)

	// 类型断言
	describeAnimal(dog)
	describeAnimal(cat)
}
