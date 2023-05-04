package main

import (
	"fmt"
	"reflect"
)

/*
	反射获取函数信息
*/

// 定义函数

func Add(a, b int) int {
	return a + b
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
}
