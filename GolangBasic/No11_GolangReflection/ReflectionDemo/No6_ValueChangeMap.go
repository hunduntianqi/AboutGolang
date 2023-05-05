package main

import (
	"GolangBaisc/No11_GolangReflection/ReflectionDemo/StructTestData"
	"fmt"
	"reflect"
)

/*
	通过 reflect.Value 修改 Map 数据
*/

func main() {
	//basicDataMap()
	structDataMap()
}

// 定义函数, 测试value为基本数据的Map
func basicDataMap() {
	// 创建一个Map对象
	mapTest := make(map[int]string)
	// map中存入数据
	mapTest[1] = "1"
	mapTest[2] = "2"
	mapTest[3] = "3"
	fmt.Println(mapTest)
	// 创建 指针Value对象
	mapPtrValue := reflect.ValueOf(&mapTest)
	// 修改map数据 ==> SetMapIndex(key, val Value)
	mapPtrValue.Elem().SetMapIndex(reflect.ValueOf(1), reflect.ValueOf("一"))
	mapPtrValue.Elem().SetMapIndex(reflect.ValueOf(2), reflect.ValueOf("二"))
	mapPtrValue.Elem().SetMapIndex(reflect.ValueOf(3), reflect.ValueOf("三"))
	// 添加键值对 ==> SetMapIndex(key, val Value), key之前不存在map中
	mapPtrValue.Elem().SetMapIndex(reflect.ValueOf(4), reflect.ValueOf("四"))
	fmt.Println(mapTest)
	// 获取 map 中所有 key 的 reflect.Value 切片
	keyValues := mapPtrValue.Elem().MapKeys()
	// 遍历获取所有的值
	for _, key := range keyValues {
		// reflect.Interface.(type) ==> 将任意类型的原始数据转换为特定数据类型
		fmt.Println(mapTest[key.Interface().(int)])
	}
}

// 定义函数, 测试value为结构体的Map
func structDataMap() {
	// 创建Map对象
	userMap := make(map[string]*StructTestData.UserMessage)
	// 创建 StructTestData.UserMessage 结构体对象
	user1 := StructTestData.UserMessage{
		Name:  "郭鹏涛",
		Age:   25,
		Sex:   '男',
		Phone: "17320101759",
		Email: "1729992141@qq.com",
	}
	userMap[user1.Name] = &user1
	// 创建Map 的 指针Value对象
	structMapValue := reflect.ValueOf(&userMap)
	fmt.Println(structMapValue)
	// 根据键名的 reflect.Value 获取 对应 值的 reflect.Value 对象
	mapStructValue := structMapValue.Elem().MapIndex(reflect.ValueOf(user1.Name))
	fmt.Println(mapStructValue)
	// 修改结构体字段信息
	fmt.Println("修改信息前 ==> ", user1)
	mapStructValue.Elem().FieldByName("Age").Set(reflect.ValueOf(26))
	fmt.Println("修改信息后 ==> ", user1)
	// 向map中添加数据
	user2 := StructTestData.UserMessage{
		Name:  "郭鹏强",
		Age:   22,
		Sex:   '男',
		Phone: "13598188577",
		Email: "2249045011@qq.com",
	}
	structMapValue.Elem().SetMapIndex(reflect.ValueOf(user2.Name), reflect.ValueOf(&user2))
	fmt.Println(userMap)
	// 获取 map中所有 key 的 Value 切片
	keyValues := structMapValue.Elem().MapKeys()
	for _, key := range keyValues {
		fmt.Println(userMap[key.Interface().(string)])
	}
}
