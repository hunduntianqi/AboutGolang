package main

import "fmt"

/*
	格式化输出函数
		fmt.Printf(格式化字符串, 变量)
		字符串格式化时常用动词及功能:
			%v: 按照变量的值进行输出
			%+v: 在 %v 基础上, 对结构体字段名和值进行展开
			%#v: 输出 Go 语言语法格式的值
			%T: 输出Go语言语法格式的类型和值
			%%: 输出%
			%b: 整型数据以二进制方式显示
			%o: 整形数据以八进制方式显示
			%d: 整型数据以十进制方式显示
			%x: 整型数据以十六进制方式显示
			%X: 整型以十六进制, 字母大写的方式显示
			%U: Unicode字符
			%f: 浮点数
			%p: 指针, 十六进制显示
*/

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14
	// %T操作输出变量所属类型
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)
	// 精确格式化
	fmt.Printf("%d, %s, %c, %f\n", a, b, c, d)
	// 万能格式化%v
	fmt.Printf("%v, %v, %v, %v", a, b, c, d)
}
