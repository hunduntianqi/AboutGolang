package main

import (
	"fmt"
	"reflect"
)

/*
	反射获取函数信息
*/

// 定义函数

func Add(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	//反射获取函数类型
	funcType := reflect.TypeOf(Add)
	// 获取函数形参个数
	fmt.Println("函数 Add 的形参个数为 ==> ", funcType.NumIn())
	// 循环遍历获取函数参数数据类型
	for i := 0; i < funcType.NumIn(); i++ {
		fmt.Printf("函数 Add 的第 %d 个参数的数据类型为 ==> %v\n", i+1, funcType.In(i))
	}
	fmt.Println("======================")
	// 获取函数返回值个数
	fmt.Println("函数 Add 的返回值个数为 ==> ", funcType.NumOut())
	// 循环遍历获取函数返回值数据类型
	for i := 0; i < funcType.NumOut(); i++ {
		fmt.Printf("函数 Add 的第 %d 个参数的返回值数据类型为 ==> %v\n", i+1, funcType.Out(i))
	}
	// 调用函数
	// 创建函数的Value对象
	funcValue := reflect.ValueOf(Add)
	// 创建存储函数参数的切片
	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf(2), reflect.ValueOf(2))
	fmt.Println(funcValue.Call(args)[0].Interface().(int))
	fmt.Println(funcValue.Call(args)[1].Interface().(int))
}
