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
	structTest()
	fmt.Println("=============================")
	basicDataTest()
}

// 定义函数, 测试结构体相关数据
func structTest() {
	// 创建结构体对象
	user := StructTestData.UserMessage{
		Name:  "郭鹏涛",
		Age:   25,
		Sex:   '男',
		Phone: "17320101759",
		Email: "1729992141@qq.com",
	}
	// 获取指针 reflect.Value 对象
	structPtrValue := reflect.ValueOf(&user)
	// 获取变量的数据类型
	fmt.Println("变量 structPtrValue 的完整数据类型为 ==> ", structPtrValue.String())
	fmt.Println("变量 structPtrValue 的 Kind 数据类型为 ==> ", structPtrValue.Kind())
	// Value 对象转换为 Type 类型对象
	structType := structPtrValue.Type()
	fmt.Println("将 structPtrValue 的 reflect.Value 转换为 reflect.Type 后的实际数据类型为 ==> ", reflect.TypeOf(structType))
	// 通过 reflect.Value.Interface() 获取 Value 保存变量的原始数据
	userData := structPtrValue.Interface()
	fmt.Println("structPtrValue 保存的原始数据为 ==> ", userData)
	// 指针 reflect.Value 对象 转换为非指针 reflect.Value 对象
	structValue := structPtrValue.Elem()
	fmt.Println("变量 structValue 的 reflect.Value 完整数据类型为 ==> ", structValue.String())
	fmt.Println("变量 structValue 的 reflect.Value.Kind() 数据类型为 ==> ", structValue.Kind())
	fmt.Println("structValue 保存的原始数据为 ==> ", structValue.Interface())

	// 通过 reflect.Value 修改原始数据的值
	fmt.Println("user 被修改前的原始数据为: ", structValue.Interface())
	structPtrValue.Elem().FieldByName("Name").Set(reflect.ValueOf("郭鹏强"))
	structPtrValue.Elem().FieldByName("Age").Set(reflect.ValueOf(22))
	structPtrValue.Elem().FieldByName("Phone").Set(reflect.ValueOf("13598188577"))
	structPtrValue.Elem().FieldByName("Email").Set(reflect.ValueOf("2249045011@qq.com"))
	fmt.Println("user 被修改后的原始数据为: ", structValue.Interface())
}

// 定义函数, 测试基本数据类型
func basicDataTest() {
	// 定义基本数据变量
	var intData int = 10
	var strData string = "测试数据"
	var floatData float64 = 64.5

	// 创建基本数据类型的 reflect.Value 对象
	intValue := reflect.ValueOf(intData)
	fmt.Println("变量 intValue 的完整数据类型为 ==> ", intValue.String())
	fmt.Println("变量 intValue 的 Kind 数据类型为 ==> ", intValue.Kind())

	strValue := reflect.ValueOf(strData)
	fmt.Println("变量 strValue 的完整数据类型为 ==> ", strValue.String())
	fmt.Println("变量 strValue 的 Kind 数据类型为 ==> ", strValue.Kind())

	floatValue := reflect.ValueOf(floatData)
	fmt.Println("变量 floatValue 的完整数据类型为 ==> ", floatValue.String())
	fmt.Println("变量 floatValue 的 Kind 数据类型为 ==> ", floatValue.Kind())

	fmt.Println()
	// 创建基本数据类型的 指针 reflect.Value 对象
	intPtrValue := reflect.ValueOf(&intData)
	fmt.Println("变量 intPtrValue 的完整数据类型为 ==> ", intPtrValue.String())
	fmt.Println("变量 intPtrValue 的 Kind 数据类型为 ==> ", intPtrValue.Kind())

	strPtrValue := reflect.ValueOf(&strData)
	fmt.Println("变量 strPtrValue 的完整数据类型为 ==> ", strPtrValue.String())
	fmt.Println("变量 strPtrValue 的 Kind 数据类型为 ==> ", strPtrValue.Kind())

	floatPtrValue := reflect.ValueOf(&floatData)
	fmt.Println("变量 floatValue 的完整数据类型为 ==> ", floatPtrValue.String())
	fmt.Println("变量 floatValue 的 Kind 数据类型为 ==> ", floatPtrValue.Kind())
	fmt.Println()
	// 通过 Value 修改原始数据
	fmt.Println("intData被修改前的数据为: ", intData)
	intPtrValue.Elem().SetInt(100)
	fmt.Println("intData被修改后的数据为: ", intData)
	fmt.Println("strData被修改前的数据为: ", strData)
	strPtrValue.Elem().SetString("修改测试数据")
	fmt.Println("strData被修改后的数据为: ", strData)
	fmt.Println("floatData被修改前的数据为: ", floatData)
	floatPtrValue.Elem().SetFloat(81.81)
	fmt.Println("floatData被修改后的数据为: ", floatData)
}
