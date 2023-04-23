package main

import "fmt"

/*
	变量:
		变量声明:
			var 变量名 变量类型
			例:
				var name string
				var age int
		批量声明:
			var (
				变量名1 变量1的类型
				变量名2 变量2的类型
			)
		变量的默认值:
			整型和浮点型: 0
			字符串: ""
			布尔型: false
			切片, 函数, 指针变量:nil
		声明变量时初始化:
			var 变量名 变量类型 = 变量值
			例:
				var name string = "pprof.cn"
				var sex int = 1
			同时初始化多个变量:
				var 变量1 变量2 = 变量1的值, 变量2的值
			在函数内部, 可以使用更简略的 := 方式声明并初始化变量:
				变量名 := 变量值
		匿名变量:
			匿名变量用一个下划线_表示, 在使用多重赋值时, 如果想要忽略某个值, 可以使用下划线代表匿名变量
			变量1, _ := 值1, 值2
			值1会赋值给变量1, 值2被忽略
			匿名变量不占用命名空间, 不会分配内存, 所以匿名变量之间不存在重复声明
		变量注意事项:
			1. 函数外的每个语句都必须以关键字开始(var, const, func等)
			2. := 不能用在函数外
			3. "_"多用于占位, 表示忽略值
		变量的类型转换:
			Go语言中不存在隐士转换, 类型转换必须是显示的, 并且类型转换只能发生在两种兼容类型之间
			类型转换格式: 变量名称 := 目标类型(需要转换的变量)
		变量可见性规则:
			1. 公有变量: 大写字母开头的变量是可导出的, 指其他包可以读取的变量
			2. 私有变量: 小写字母开头的, 不可导出的变量
		查看变量类型:reflect.TypeOf(变量名称)
		变量作用域:
	        作用域: 指已声明标识符所表示的常量, 类型, 变量, 函数或包在源代码中的作用范围
	        局部变量: 函数内定义的变量
	            作用域: 只在函数体内, 参数和返回值变量也是局部变量
	        全局变量: 函数外定义的变量
	            作用域: 可以在任何函数中使用, 也可以在整个包甚至外部包(被导出后)使用
				注意: 当全局变量和局部变量名称相同时, 会优先使用函数内的局部变量
	        形式参数: 函数定义中的变量
	            作用域: 会作为函数的局部变量使用
*/

func main() {
	varDefine()
}

// 变量声明实例
func varDefine() {
	// 定义单个变量并赋值
	var str string = "这是一次变量定义的尝试！！"
	fmt.Println(str)
	// 同时定义多个变量并赋值
	var c, b int = 100, 50
	fmt.Println("c = ", c, "\nb = ", b)
}