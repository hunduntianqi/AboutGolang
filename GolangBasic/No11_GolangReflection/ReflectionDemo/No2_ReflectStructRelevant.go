package main

import (
	"fmt"
	"reflect"
)

/*
	反射中与结构体相关的函数与方法
*/

// 定义一个结构体

type UserMessage struct {
	name  string `name:"姓名"`
	age   int    `age:"年龄"`
	sex   rune   `sex:"姓名"`
	phone string `phone:"联系方式"`
	email string `email:"邮箱地址"`
	int   `int:"测试用匿名字段"`
}

// 结构体绑定方法, 打印结构体对象信息
func (user UserMessage) printMessage() {
	fmt.Printf("name:%s, age:%d, sex:%r, phone:%s, email:%s", user.name, user.age, user.sex, user.phone, user.email)
	fmt.Println()
}

// 结构体绑定方法, 修改name属性信息
func (user UserMessage) setName(name string) {
	user.name = name
}

// 结构体绑定方法, 修改age属性信息
func (user UserMessage) setAge(age int) {
	user.age = age
}

// 结构体绑定方法, 修改phone属性信息
func (user UserMessage) setPhone(phone string) {
	user.phone = phone
}

func main() {
	// 获取结构体类型对应的reflect.Type对象
	structType := reflect.TypeOf(UserMessage{
		name:  "郭鹏涛",
		age:   25,
		sex:   '男',
		phone: "17320101759",
		email: "1729992141@qq.com",
	})
	// 调用函数, 显示结构体字段信息
	getFieldMessage(structType)
	// 调用函数, 显示结构体成员方法信息
	getMethodMessage(structType)
}

// 定义函数, 获取结构体字段信息
func getFieldMessage(structType reflect.Type) {
	// 获取结构体字段个数
	fmt.Println("结构体 UserMessage 中字段的个数为 ==> ", structType.NumField())
	// 获取结构体字段属性信息
	for i := 0; i < structType.NumField(); i++ {
		fmt.Println("=================================")
		// 根据索引获取字段对象
		field := structType.Field(i)
		fmt.Println("字段名称 ==> ", field.Name)
		fmt.Println("相对于首地址内存偏移量 ==> ", field.Offset)
		fmt.Println("是否为匿名字段 ==> ", field.Anonymous)
		fmt.Println("字段数据类型 ==> ", field.Type)
		fmt.Println("字段的标签 ==> ", field.Tag)
		fmt.Println("字段非导出包路径 ==> ", field.PkgPath)
	}
}

// 定义函数, 获取结构体成员方法信息
func getMethodMessage(structType reflect.Type) {
	// 获取结构体成员方法的个数
	fmt.Println("结构体 UserMessage 的成员方法个数为 ==> ", structType.NumMethod())
	// 根据索引获取对应成员方法
	structType.Method(0)
}
