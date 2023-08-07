package main

import "fmt"

type notifier interface {
	notify()
}

type myUser struct {
	name  string
	email string
}

type myAdmin struct {
	name  string
	email string
}

type superAdmin struct {
	myUser
	level string
}

// 指针接收者定义的多态方法
// 作为 notifier 接口定义的参数只有指针传递才可以以 notifier 的身份调用 notify
// 如果是以值传递, 那么无论是指针传递还是值传递都可以
func (u *myUser) notify() {
	fmt.Println("user, name: " + u.name + ", email: " + u.email)
}

func (u *myAdmin) notify() {
	fmt.Println("admin, name: " + u.name + ", email: " + u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := myUser{"bill", "bill@mail.com"}
	// cannot use u (variable of type myUser) as notifier value in argument to sendNotification:
	// myUser does not implement notifier (method notify has pointer receiver)
	// 这里因为是值传递给指针接收者, 所以编译失败
	// sendNotification(u)

	// 必须是指针传递给指针接收者才可以
	sendNotification(&u)

	a := myAdmin{"bill", "bill@mail.com"}

	sendNotification(&a)

	sa := superAdmin{
		myUser{"superAdmin", "super@mail.com"},
		"super",
	}

	// type superAdmin struct {
	//	myUser
	//	level string
	//}
	// superAdmin 没有实现 notifier 接口, 但是也可以调用 notify 函数
	// 编译器会认为这个指针实现了 notifier 接口, 并接受值传递
	// 也可以通过 superAdmin 实现 notifier 接口, 实现 superAdmin 的逻辑
	sendNotification(&sa)
}
