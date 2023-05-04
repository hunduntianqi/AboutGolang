package main

import (
	"GolangBaisc/No11_GolangReflection/ReflectionDemo/StructTestData"
	"fmt"
	"reflect"
)

/*
	通过 reflect.TypeOf(variableName) 获取 reflect.Type对象
*/

func main() {
	// 简单获取 reflect.Type 类型变量
	typeI := reflect.TypeOf(1)
	typeS := reflect.TypeOf("hello")
	fmt.Println(typeI) //int
	fmt.Println(typeS) //string

	// 创建结构体指针的 reflect.Type 对象
	typeUser := reflect.TypeOf(&StructTestData.UserMessage{})
	fmt.Println("typeUser的完整数据类型为 ==> ", typeUser.String())
	// ptr ==> 表示typeUser的数据类型是一个指针
	fmt.Println("typeUser的 Kind 类型为 ==> ", typeUser.Kind())
	//  reflect.Type.Elem() ==> 将指针Type转换为非指针Type
	fmt.Println("typeUser的数据类型转换为非指针Type ==> typeUser.Elem(): ", typeUser.Elem())

	// 创建一个结构体 reflect.Type 对象
	typeUser2 := reflect.TypeOf(StructTestData.UserMessage{})
	fmt.Println("typeUser2的完整数据类型为 ==> ", typeUser2.String())
	fmt.Println("typeUser2的Kind数据类型为 ==> ", typeUser2.Kind())
	// 非指针Type无法转换为指针Type, 如下代码会报错: panic: reflect: Elem of invalid type StructTestData.UserMessage
	// fmt.Println("typeUser2 转换为指针 Type 数据类型为 ==> ", typeUser2.Elem())
}
