package domain

// User 是领域模型（Service 层使用的业务对象）。
// 注意：密码相关字段不对外输出；后续可根据需要调整为更细粒度的模型。
type User struct {
	ID       uint
	Username string
	Email    string
}

