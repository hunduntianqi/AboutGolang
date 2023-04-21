package main

import "fmt"

/*
   方法集:
        在Go语言中, 每个类型都有与之相关的方法, 这个类型的所有方法成为类型的方法集
        1. 类型 T 方法集, 包含全部的receive T 方法
        2. 类型 *T 方法集, 包含全部的receive T + *T 方法
        3. 如类型 S 包含匿名字段 T, 则 S 和 *S 方法集包含 T 方法
        4. 如类型 S 包含匿名字段 *T, 则 S 和 *S 方法集包含 T + *T 方法
        5. 不管嵌入 T 或 *T, *S 方法集总是包含 T + *T 方法
        使用实例的 value 和 pointer 调用方法(含匿名字段), 不受方法集的约束, 编译器总是查找全部方法, 并自动转化receive实参
*/

// 定义类型T
type T struct {
	int
}

// 定义包含匿名字段T的类型S
type S struct {
	T
}

// 定义类型T的方法
func (t T) methodValueT() {
	fmt.Println("类型 T 方法集包含全部receive T 的方法")
}

// 定义类型 *T 的方法
func (t T) methodPointT1() {
	fmt.Println("类型 *T 方法集包含全部receive T 的方法")
}

func (t *T) methodPointT2() {
	fmt.Println("类型 *T 方法集包含全部receive *T 的方法")
}

// 定义类型S的方法
func (t T) sMethod() {
	fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法")
}
func main() {
	// 定义类型T变量
	t := T{}
	t.methodValueT()
	// 定义类型*T变量
	t2 := &t
	t2.methodPointT1()
	t2.methodPointT2()

	// 定义类型S变量
	s := S{T{1}}
	s.sMethod()
	// 定义类型*S变量
	s2 := &s
	s2.sMethod()
}
