package main

import "fmt"

/*
	表达式:
		根据调用者不同, 方法分为两种表现形式
		method value ==> instance.method(args...):
		method expression ==> <type>.func(instance, args...)
		两者都可像普通函数那样赋值和传参, 区别在于 method value 绑定实例, 而 method expression 则须显式传参
		注意:
			1. method value 会复制 receiver, 由于不是指针类型, 后续receiver一旦修改, 不会改变method value中的receiver内容
			  在汇编层面, method value 和闭包的实现方式相同, 实际返回 FuncVal 类型对象
			  ==>FuncVal { method_address, receiver_copy }
*/

func main() {
	methodValueTest()
}

// 定义User结构体
type User struct {
	id   int
	name string
}

// 为User结构体定义方法
func (self User) printMessage() {
	fmt.Printf("id = %d, name = %s\n", self.id, self.name)
}

// 定义测试函数, 演示method value和method expression两种不同形式
func methodValueAndExpression() {
	// 定义User类型变量
	user := User{
		id:   1,
		name: "混沌天炁",
	}
	user.printMessage()
	mValue := user.printMessage // 隐式传递receiver
	mValue()
	mExpression := (*User).printMessage
	mExpression(&user) // 显示传递receiver
}

// 定义方法, 打印User实例接收者
func (self User) printSelf() {
	fmt.Println(self)
}

// method value调用形式会复制receiver
func methodValueTest() {
	// 定义User类型变量
	user := User{1, "混沌天炁"}
	// 该调用方法会将receiver user复制给mValue变量, 由于不是指针, 后续user对象成员修改时, 不会影响到mValue中receiver的内容
	mValue := user.printSelf
	user.name = "造化轮回"
	user.printSelf() // {1 造化轮回}
	mValue()         // {1 混沌天炁}
}
