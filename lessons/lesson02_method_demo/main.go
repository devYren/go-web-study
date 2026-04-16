package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 属于 User | 入参: greeting string | 出参: string
func (u *User) SayHello(greeting string) (string, error) {
	return fmt.Sprintf("%s! 我叫%s, 今年%d岁", greeting, u.Name, u.Age), nil
}

func main() {
	u := User{Name: "Alice", Age: 25}

	// u → 接收者（属于谁）
	// "你好" → 入参
	// msg → 接收出参
	msg, _ := u.SayHello("你好")
	fmt.Println(msg)
}
