package main

import "fmt"

/*
	数据输入:
		输入函数:
			fmt.Scanf("格式化字符串", &接收键盘输入信息的变量)
			fmt.Scan(&接收键盘输入信息的变量)
*/

func main() {
	var a string
	fmt.Println("请输入变量a的值:")
	// 会阻塞等待用户输入
	fmt.Scan(&a)
	fmt.Println("变量a的值为:", a)

	// 定义变量接收键盘输入数据
	fmt.Println("请输入你要说的话:")
	var text string
	fmt.Scan(&text)
	fmt.Println(text)
}
