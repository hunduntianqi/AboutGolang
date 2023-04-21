/*
	defer关键字作用:
		用于延迟调用一个函数或者方法(或者当前所创建的匿名函数)
		defer用途:
			1. 关闭文件句柄
			2. 锁资源释放
			3. 数据库连接释放
		注意:defer语句只能出现在函数或方法的内部
		使用格式:
			函数名 func() {
				代码1执行
				代码2执行
				defer 代码3执行 //此语句会在函数执行结束前调用, 函数会先执行此语句后面的代码3
				代码3执行
			}
		多个的defer语句执行顺序:
			如果一个函数中有多个defer语句, 会以LIFO(后进先出, 从下往上)的顺序执行, 即使函数或某个延迟调用发生错误, 这些调用
			仍然会被执行
		defer语句和匿名函数结合执行:
			defer func() {
				功能代码
			} ()
			匿名函数的执行在其他defer修饰代码之前
		defer先进后出原则: 后面的语句会依赖前面的资源, 因此如果先前面的资源先释放了, 后面的语句就没法执行了
*/

package main

import "fmt"

func main() {
	deferDemo1()
	deferDemo2()
}

// 定义函数测试defer先进后出原则
func deferDemo1() {
	var whatever [5]int
	for index, _ := range whatever {
		defer fmt.Println(index) // 因为defer先进后出原则, 会先打印数组最后一个元素, 再从后往前依次打印数组元素
	}
}

// defer 遇上闭包
func deferDemo2() {
	var whatever [5]int
	for index, _ := range whatever {
		// 由于闭包对外部函数变量是引用传递, 所以defer修饰的闭包在延迟调用时, 输出的外部函数变量内容一样
		defer func() { fmt.Println(index) }()
	}
}
