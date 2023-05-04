package main

import (
	"GolangBaisc/No11_GolangReflection/ReflectionDemo/StructTestData"
	"fmt"
	"reflect"
)

/*
	通过 reflect.ValueOf(variableName) 获取 reflect.Value对象
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
	// 获取 reflect.Value 对象
	structValue := reflect.ValueOf(user)
	// 获取变量的数据类型
	fmt.Println(structValue.Type())
}
