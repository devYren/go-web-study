package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func findUser(id int) *User {
	if id == 1 {
		return &User{Name: "Alice", Age: 25}
	}
	return nil // 没找到，返回 nil（对标 Java 返回 null）
}

func main() {
	// ==========================================
	// 1. nil 指针 — 对标 Java 的 null
	// ==========================================
	var p *User    // 声明了指针但没赋值，默认是 nil
	fmt.Println(p) // <nil>

	// 访问 nil 指针会 panic（对标 Java 的 NullPointerException）
	// fmt.Println(p.Name) // ❌ panic: runtime error: invalid memory address

	// 所以使用指针前要判 nil（对标 Java 判 null）
	if p != nil {
		fmt.Println(p.Name)
	} else {
		fmt.Println("p 是 nil")
	}

	// ==========================================
	// 2. 实际用法：函数返回 nil 表示"没找到"
	// ==========================================
	user := findUser(1)
	if user != nil {
		fmt.Println("找到了:", user.Name)
	}

	user2 := findUser(999)
	if user2 == nil {
		fmt.Println("没找到")
	}

	// ==========================================
	// 3. 什么时候用指针？简单记忆
	// ==========================================
	//
	// 用指针（*Type）的场景：
	//   ✅ 函数需要修改传入的值      → func modify(u *User)fast

	//   ✅ struct 比较大，避免拷贝开销 → func process(data *BigStruct)
	//   ✅ 需要表达"可能为空"         → func find() *User（nil = 没找到）
	//   ✅ 方法接收者（大部分用指针）  → func (u *User) SetName(...)
	//
	// 不用指针（直接传值）的场景：
	//   ✅ 基本类型（int, string, bool）→ 本身就很小，拷贝无所谓
	//   ✅ 只读操作，不需要修改        → func (u User) String() string
	//   ✅ 小 struct（2-3 个字段）     → 拷贝开销可以忽略
	//
	// ==========================================
	// 4. 对照项目中的实际用法
	// ==========================================
	//
	// 你项目里到处都是指针，因为 Web 开发中 struct 通常需要被修改或共享：
	//
	// internal/repo/user_repo.go:
	//   func NewUserRepo(db *gorm.DB) UserRepository
	//   → db 用指针：整个应用共享一个数据库连接，不能拷贝
	//
	//   func (r *userRepo) Create(ctx context.Context, u *model.UserModel) error
	//   → r 用指针接收者：访问 r.db
	//   → u 用指针：Create 会往 u.ID 里写入自增 ID，调用方要能看到
	//
	//   func (r *userRepo) FindByID(...) (*model.UserModel, error)
	//   → 返回 *model.UserModel：可能返回 nil 表示没找到
	//
	// internal/services/impl/auth_service_impl.go:
	//   func NewAuthService(userRepo repo.UserRepository) services.AuthService
	//   → userRepo 是接口类型，接口本身就是引用语义，不需要 *

	fmt.Println("指针课程完成")
}
