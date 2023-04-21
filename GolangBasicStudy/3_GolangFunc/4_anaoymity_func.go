package main

import "fmt"

/*
	匿名函数:
       指不需要定义函数名的一种函数实现方式
       Go语言中, 函数可以向普通变量一样被传递和使用, Go语言中支持随时在代码里定义匿名函数
       匿名函数由一个不带函数名的函数声明和函数体组成, 匿名函数的优越性在于可以直接使用函数内的
       变量, 不必声明
        匿名函数定义:
            var_name := func (param_name param_type) (return_type) {
                            函数体
                            return
                        }
        匿名函数调用:
            无返回值: var_name(实参)
            有返回值: var := var_name(实参)
*/

func main() {
	// 定义无参数无返回值匿名函数
	printNonNameDemo := func() {
		fmt.Println("这是一个匿名函数~~~")
	}
	printNonNameDemo()
	// 定义带参数无返回值匿名函数
	nonNameFunc1 := func(name string) {
		fmt.Printf("我的名字是%s\n", name)
	}
	nonNameFunc1("郭鹏涛")
	// 定义带参数带返回值匿名函数
	nonNameFunc2 := func(name string) (age int) {
		fmt.Printf("%s的年龄是多少?\n", name)
		fmt.Scan(&age)
		return
	}
	age := nonNameFunc2("郭鹏涛")
	fmt.Printf("age = %d\n", age)
}
