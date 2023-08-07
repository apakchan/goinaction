package main

import "fmt"

type Duration int64

type user struct {
	id    int
	name  string
	email string
}

// 值接收者
// 将调用者复制到此函数内
// 如果调用者在这个函数内状态改变，不会影响 src 调用者
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u user) changeName(newName string) {
	u.name = newName
}

// 指针接收者
// 共用同一个地址
func (u *user) changeEmail(newEmail string) {
	u.email = newEmail
}

type admin struct {
	person user
	level  string
}

func createPerson() {
	var bill user
	fmt.Println(bill) // {0  }
	_ = user{
		id:    1,
		name:  "lise",
		email: "abc@def.com",
	}
	_ = user{2, "lisa", "asd@asd.com"}

	_ = admin{
		person: user{1000, "fred", "asd@qq.com"},
		level:  "super",
	}
}

func testStructMethod() {
	// 值
	bill := user{1, "bill", "@qq.com"}
	// 指针
	lisa := &user{2, "lisa", "@qq.com"}

	// 值和指针都可以传递给值接收者声明的方法
	bill.notify()
	lisa.notify()

	// 值和指针都可以传递给指针接收者声明的方法
	bill.changeEmail("@asd.com")
	lisa.changeEmail("@asd.com")

	bill.notify()
	lisa.notify()

	bill.changeName("newBill")
	lisa.changeName("newLisa")

	bill.notify() // Sending User Email To bill<@asd.com>
	lisa.notify() // Sending User Email To lisa<@asd.com>
}

func main() {
	testStructMethod()
}
