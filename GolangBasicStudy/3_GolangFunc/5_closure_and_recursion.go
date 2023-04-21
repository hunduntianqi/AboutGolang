package main

import "fmt"

/*
    闭包:
		指一个函数捕获了和它在同一作用域的其他常量和变量,当闭包被调用时,不管在程序什么地方调用, 闭包都可以使用这些常量或者变量
		闭包不关心捕获的常量和变量是否已经超出了作用域, 所以只要闭包还在使用这些常量和变量,这些常量和变量就会存在
        在Go语言中,所有的匿名函数都是闭包！！
        闭包捕获外部变量的特点:
            闭包(匿名函数)以引用方式捕获外部变量, 在闭包中对外部变量进行修改, 外部变量的值会同步改变
        闭包定义格式:
            func outFunc_name () {
                // 定义外部变量
                i := 0
                // 定义匿名函数
                innerFunc := func () {
                    // 在匿名函数中对外部变量进行修改
                    i++
                }
            return innerFunc // 在汇编层,return实际返回的是FuncVal对象, 其中包含匿名函数地址, 闭包对象指针
        }
	递归函数:
		递归 ==> 指在运行的过程中自己调用自己, 一个函数调用自己, 就叫做递归函数
		构成递归的条件:
			1. 子问题须与原始问题为同样的事, 且更为简单
			2. 必须设置递归结束条件, 不能无限制的递归下去
*/

func main() {
	// 调用闭包函数
	testVar := outFunc() // outFunc() 函数调用后, 返回值为innerFunc, 因此testVar实际也是函数
	testVar()
	// 闭包引用外部函数参数
	temp := add_(100)
	fmt.Println(temp(10), temp(20))
	// 接受两个闭包
	add, sub := test(100)
	fmt.Println(add(10), sub(20))
	// 调用递归函数
	numFact := numFactorial(10)
	fmt.Printf("10的阶乘是%d\n", numFact)
	// for循环结合递归函数完成斐波那契数列
	for i := 0; i < 10; i++ {
		fmt.Printf("FeiBoNaQi(%d) = %d\n", i, FeiBoNaQi(i))
	}

}

// 定义一个简单的闭包, 并作为返回值返回
func outFunc() func() {
	i := 0
	fmt.Println(&i)
	innerFunc := func() {
		i++
		fmt.Println(&i)
		fmt.Println(i)
	}
	return innerFunc
}

// 闭包引用外部函数参数
func add_(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

// 定义两个匿名函数, 返回两个闭包
func test(base int) (func(int) int, func(int) int) {
	// 加
	add := func(num1 int) int {
		base += num1
		return base
	}
	// 减
	sub := func(num2 int) int {
		base -= num2
		return base
	}
	return add, sub
}

// 定义递归函数完成数字阶乘
func numFactorial(num int) int {
	if num == 1 {
		fmt.Println(num)
		return num
	}
	return num * numFactorial(num-1)
}

// 定义递归函数完成斐波那契数列
func FeiBoNaQi(num int) int {
	// 斐波那契数列:从第三项开始, 每一项等于前两项之和
	if num == 0 {
		return num
	} else if num == 1 {
		return num
	}
	return FeiBoNaQi(num-1) + FeiBoNaQi(num-2)
}
