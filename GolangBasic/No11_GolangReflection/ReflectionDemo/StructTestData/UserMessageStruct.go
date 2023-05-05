package StructTestData

import "fmt"

// 定义一个结构体

type UserMessage struct {
	Name  string `name:"姓名"`
	Age   int    `age:"年龄"`
	Sex   rune   `sex:"姓名"`
	Phone string `phone:"联系方式"`
	Email string `email:"邮箱地址"`
	int   `int:"测试用匿名字段"`
}

// PrintMessage 结构体绑定方法, 打印结构体对象信息
func (user UserMessage) PrintMessage() {
	fmt.Printf("name:%s, age:%d, sex:%s, phone:%s, email:%s", user.Name, user.Age, string(user.Sex), user.Phone, user.Email)
	fmt.Println()
}

// SetName 结构体绑定方法, 修改name属性信息
func (user *UserMessage) SetName(name string) {
	user.Name = name
}

// SetAge 结构体绑定方法, 修改age属性信息
func (user *UserMessage) SetAge(age int) {
	user.Age = age
}

// SetPhone 结构体绑定方法, 修改phone属性信息
func (user *UserMessage) SetPhone(phone string) {
	user.Phone = phone
}
