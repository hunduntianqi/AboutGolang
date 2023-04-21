package main

import "fmt"

/*
	封装实现 ==> 方法:
		体现面向对象的封装特性(个人认为类似Java的成员方法, 与之相比, Java的成员方法只能由对象调用, 而Go语言中的方法可以与任意类型数据
		绑定)
		本质上一个方法是一个作用于特定 类型 变量 的 函数的函数, 这种特定 类型 变量叫做 接收者(Receiver),接收者的概念类似于Java中的
		this和Python中的self
		在Go语言中, 可以给任意自定义类型(包括内置类型)添加相应的方法
		方法定义格式:
			func (receiver ReceiverType) funcName(parameters) (results) {
				功能代码
			}
			注意:
				1. 参数receiver可以任意命名, 如果方法中未使用, 可以省略不写
				2. 不支持重载方法, 不能定义名字相同但是参数不同的的方法
			例:
				type int_test int
				// 给int_test类型绑定一个函数
				func (num int_test) add(b int_test) int_test {
					return num + b
				}
				func main() {
					var num int_test
					num = 10
					fmt.Println(num.add(100))
				}
			方法的调用格式:
				var_name.method(参数)
				还可以使用存储了该类型变量地址的指针调用方法
				即: point = &var_name
					point.method(参数)
		普通函数与方法的区别:
			普通函数: 接收者为值类型时, 不能将指针类型的数据直接传递, 反之亦然
			方法: 接收者为值类型时, 可以直接用指针类型的变量调用方法, 反之亦然
*/

func main() {
	// 定义User类型变量
	var user UserMethod
	// 结构体变量初始化
	user.user = "混沌天炁"
	user.email = "1729992141@qq.com"
	// 调用User类型绑定的方法
	user.userFunc()
	// 定义结构体类型的指针变量
	var userPoint *UserMethod
	userPoint = &user
	// 使用指针调用User的方法
	userPoint.userFunc()
	// 调用测试方法
	funcTest()
	methodTestMethod()
}

// 定义一个结构体类型
type UserMethod struct {
	user  string
	email string
}

// 给定义的User结构体定义方法
func (u UserMethod) userFunc() {
	fmt.Printf("user = %s, email = %s\n", u.user, u.email)
}

// 普通函数
// 接收者为值类型
func valueFunc(num int) int {
	return num + 10
}

// 接收者为指针类型函数
func pointFunc(num *int) int {
	return *num + 10
}

// 定义普通函数测试函数
func funcTest() {
	// 定义普通值类型变量
	a := 2
	fmt.Println("valueFunc(2) = ", valueFunc(a))
	// fmt.Println("valueFunc(2) = ", valueFunc(&a)) ==> 接收者为值类型函数, 不能直接传入地址(指针)参数
	fmt.Println("pointFunc(2) = ", pointFunc(&a))
	// fmt.Println("pointFunc(2) = ", pointFunc(a)) ==> 接收者为指针类型函数, 不能直接传入值类型参数
}

// 方法
// 接收者为值类型方法
func (user UserMethod) valueMethod() {
	fmt.Printf("user = %s, email = %s\n", user.user, user.email)
}

// 接收者为指针类型方法
func (user *UserMethod) pointMethod() {
	fmt.Printf("user = %s, email = %s\n", user.user, user.email)
}

// 定义方法测试函数
func methodTestMethod() {
	// 定义普通User类型变量
	user := UserMethod{user: "混沌天炁", email: "1729992141@qq.com"}
	user.valueMethod() // 可以调用值类型接收者方法
	user.pointMethod() // 可以调用指针类型接收者方法
	// 定义指针类型User变量
	var user2 *UserMethod
	user2 = &user
	user2.valueMethod() // 可以调用值类型接收者方法
	user2.pointMethod() // 可以调用指针类型接收者方法
}
