package main

import "fmt"

/*
	常量:
		常量是恒定不变的值, 多用于定义程序运行期间不会改变的那些值, 常量在定义时必须赋值
		常量定义形式:
			显示定义:const 常量名称 常量类型 = 常量值
			隐式定义: const 常量名称 = 常量值(通常叫无类型常量)
			组合定义:
			const (
				常量名 常量类型 = 常量值
				常量名 = 常量值
			)
			定义多个常量:
			const 常量1, 常量2, ...常量类型 = 常量值1, 常量值2, ...
			注意:
				1.常量可以使用内置表达式定义, 例如: len(), unsafe.Sizeof()等
		常量支持类型:常量目前只支持布尔型, 数字型, 和字符串型
		iota:
			是Go语言的常量计数器, iota的值在const关键字出现时会被重置为0,
			const中每新增一行常量声明, iota自增1
			iota常见使用方法:
				1.跳值使用法
				2.插队使用法
				3.表达式隐式使用法
				4.单行使用法
*/

func main() {
	constDefine()
}

// 常量声明实例
func constDefine() {
	// 单个常量定义
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	// 多个常量同时定义
	const a, b, c = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("面积为: %d平方米\n", area)
	fmt.Println(a, b, c)
}
