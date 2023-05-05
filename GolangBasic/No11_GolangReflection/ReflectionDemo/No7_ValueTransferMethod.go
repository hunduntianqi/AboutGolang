package main

import (
	"GolangBaisc/No11_GolangReflection/ReflectionDemo/StructTestData"
	"reflect"
)

/*
	reflect.Value 调用成员方法
*/

func main() {
	// 创建结构体对象
	user := StructTestData.UserMessage{
		Name:  "郭鹏涛",
		Age:   25,
		Sex:   '男',
		Phone: "17320101759",
		Email: "1729992141@qq.com",
	}
	// 创建结构体的 reflect.Value 对象
	userValue := reflect.ValueOf(&user)
	// 根据方法名获取无参成员方法对象
	methodPrintMessage := userValue.MethodByName("PrintMessage")
	// 调用成员方法, 无参方法传入一个空切片
	methodPrintMessage.Call([]reflect.Value{})
	// 根据方法名获取带参成员方法对象
	methodSetName := userValue.MethodByName("SetName")
	// 定义切片, 存储方法参数
	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf("郭鹏强"))
	// 调用带参成员方法, 有返回值时需要变量来接收, 返回值为 []reflect.Value 类型
	methodSetName.Call(args)
	user.PrintMessage()
	args = append(args, reflect.ValueOf("郭鹏涛"))
	args = append(args[:0], args[1:]...) // 删除之前存入的第一个元素
	// 调用带参成员方法, 有返回值时需要变量来接收, 返回值为 []reflect.Value 类型
	methodSetName.Call(args)
	user.PrintMessage()
}
