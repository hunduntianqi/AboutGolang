package main

import (
	"GolangBaisc/No11_GolangReflection/ReflectionDemo/StructTestData"
	"fmt"
	"reflect"
)

/*
	通过 reflect.Value 修改切片数据
*/

func main() {
	basicDataSlice()
	//structSlice()
}

// 定义函数, 测试基本数据类型切片数据修改情况
func basicDataSlice() {
	// 定义一个切片, 可存储任意数据类型
	list := make([]interface{}, 0, 10)
	// 打印切片初始信息
	fmt.Println("list的容量为 ==> ", cap(list))
	fmt.Println("list的元素数量为 ==> ", len(list))
	// 向切片中存入数据
	list = append(list, 9, "hello world")
	// 根据切片, 创建 指针Value 对象
	listPtrValue := reflect.ValueOf(&list)
	fmt.Println("list原始容量为 ==> ", listPtrValue.Elem().Cap())
	// 修改原始切片的容量
	listPtrValue.Elem().SetCap(9)
	fmt.Println("list原始容量被修改后为 ==> ", listPtrValue.Elem().Cap())
	fmt.Println("list切片长度为 ==> ", listPtrValue.Elem().Len())
	// 修改切片长度
	listPtrValue.Elem().SetLen(4)
	fmt.Println("list切片长度被修改后为 ==> ", listPtrValue.Elem().Len())
	fmt.Println("list的元素信息 ==> ", listPtrValue.Elem().Interface())
	// 根据索引获取切片中的第四个和第五个元素并修改数据
	listPtrValue.Elem().Index(2).Set(reflect.ValueOf("切片的第3个元素"))
	listPtrValue.Elem().Index(3).Set(reflect.ValueOf("切片的第4个元素"))
	fmt.Println("list的元素信息 ==> ", listPtrValue.Elem().Interface())
}

// 定义函数, 测试结构体类型切片数据修改情况
func structSlice() {
	// 定义一个切片, 可以存储 *StructTestData.UserMessage 类型的数据
	list := make([]*StructTestData.UserMessage, 0)
	// 定义一个结构体对象
	user := StructTestData.UserMessage{
		Name:  "郭鹏涛",
		Age:   25,
		Sex:   '男',
		Phone: "17320101759",
		Email: "1729992141@qq.com",
	}
	// 将 user 的指针存入切片
	list = append(list, &user)
	// 创建 list 的指针Value对象
	listPtrValue := reflect.ValueOf(&list)
	fmt.Println("list中的元素信息为 ==> ", listPtrValue.Elem().Index(0).Elem().Interface())
	// 修改切片中结构体元素信息
	listPtrValue.Elem().Index(0).Elem().FieldByName("Name").Set(reflect.ValueOf("陈欣妮"))
	listPtrValue.Elem().Index(0).Elem().FieldByName("Age").Set(reflect.ValueOf(25))
	listPtrValue.Elem().Index(0).Elem().FieldByName("Sex").Set(reflect.ValueOf('女'))
	listPtrValue.Elem().Index(0).Elem().FieldByName("Phone").Set(reflect.ValueOf("13480194858"))
	listPtrValue.Elem().Index(0).Elem().FieldByName("Email").Set(reflect.ValueOf("564036589@qq.com"))
	fmt.Println("list中的元素信息为 ==> ", listPtrValue.Elem().Index(0).Elem().Interface())
}
