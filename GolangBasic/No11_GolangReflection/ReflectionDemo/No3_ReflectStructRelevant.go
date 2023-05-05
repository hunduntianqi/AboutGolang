package main

import (
	"GolangBaisc/No11_GolangReflection/ReflectionDemo/StructTestData"
	"fmt"
	"reflect"
)

/*
	反射中与结构体相关的函数与方法
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
	// 获取结构体类型对应的reflect.Type对象
	structType := reflect.TypeOf(user)
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
	// 指定索引获取成员方法
	methodIndex := structType.Method(0)
	fmt.Println("结构体 UserMessage 的第一个成员方法是: ", methodIndex.Name)
	// 指定方法名称获取成员方法
	methodName, ok := structType.MethodByName("PrintMessage")
	fmt.Println("获取结构体 UserMessage 的成员方法结果:", ok, "; methodType ==> ", methodName.Type)
}
