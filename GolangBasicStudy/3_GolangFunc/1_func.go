package main

import "fmt"

/*
	函数:
		是Go语言的基本代码块
        函数定义:
            func 函数名 (形参) (返回值类型) {
                函数体
				return value
            }
            func: 函数声明由func关键字定义
			注意: 函数如果定义有返回值类型, 则必须有数据返回, 否则会产生编译错误
        函数调用:
            无返回值:
                函数名 (传入参数)
            有返回值:
                变量 = 函数名(传入参数)
                注意: go语言中支持多返回值, 所以用来接收返回值的变量可能不止一个
		可变参数:
			即参数个数不固定, 参数类型固定的函数参数
			定义格式:
				func func_name (args ...type) (返回值类型列表) {
					函数体
					return 返回值
				}
				args: 函数参数名
				...: 代表改参数是可变参数
				type: 可变参数类型
				注意: 定义可变参数必须是形参列表的最后一个参数, 且"..."不能省略
				可变参数本质是一个slice, 参数赋值时可以通过数组或切片传参; 可以通过args[index]的形式访问可变参数, 还可以通过len(args)
				来查看可变参数中传递的参数个数, 使用slice做可变参数传参时, 必须展开 ==> slice_name...
				任意类型的可变参数: 指函数的参数和每个参数的类型是不定的, 可以是任意类型
					实现方式==>interface{}==>使用空接口定义可变参数类型==>args ...interface{}
		注意: go语言中无默认参数！！！
		函数参数传递方式:
			1. 值传递: 指在调用函数时, 只将实参的值传入函数, 在函数内部对形参的操作不会影响外部实参
			2. 引用传递: 在调用函数时, 将实参的地址值传递进函数中, 在函数内部的操作将会同步改变外部实参的值
			默认情况下, go语言使用的是值传递
			注意:
				1. 无论是值传递还是引用传递, 传递给函数的都是变量的副本, 区别在于值传递是值的拷贝, 引用传递是地址的拷贝;
					一般来说, 地址拷贝更为高效; 值拷贝取决于拷贝对象的大小, 对象越大, 则性能越低
				2. map, slice, chan, 指针, interface默认以引用方式传递
        函数返回值:
            "_"标识符: 可以用来忽略不需要的函数返回值
            go语言中可以将返回值在返回值列表中进行命名, 并向使用函数参数一样当做局部变量使用
            定义格式:
                func func_name (param_name param_type) (return_name return_type) {
                    函数体
                    return
                }
            注意:
                1. 将返回值命名后, 可以在return语句后省略返回值参数名, 该方法为隐式返回
                2. 如果在函数内又定义了与命名返回参数同名参数时, 命名返回参数会被遮蔽, 此时必须显示返回(return语句后加返回值变量名)
                3. 函数返回值只能由变量来接收, 不能使用slice等容器对象接收
                4. 多返回值可以直接作为其他函数的实参进行传递, 不需要再定义变量接收返回值再传递函数实参
                5. 命名返回参数允许 defer 延迟调用通过闭包读取和修改
*/

func main() {
	// 定义两个int变量, 调用函数输出最大值
	var num1 = 20
	var num2 = 30
	fmt.Println("num1 和 num2之间的最大值是:", max(num1, num2))
	// 调用多返回值函数
	name, age := swap("郭鹏涛", 24)
	fmt.Println("name:", name, "age:", age)
	nextNumber := getSequence()
	for i := 0; i < 3; i++ {
		fmt.Println(nextNumber())
	}
	// 定义空切片
	list := make([]interface{}, 5)
	changeParam(list, 1, 2, 3, 4, 5, '男', "郭鹏涛")
}

// 定义函数判断最值
func max(num1, num2 int) int {
	// 定义中间变量
	var temp int
	if num1 > num2 {
		temp = num1
	} else {
		temp = num2
	}
	return temp
}

// 定义多返回值函数
func swap(name string, age int) (string, int) {
	return name, age
}

// 定义闭包函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 定义可变参数函数
func changeParam(name []interface{}, args ...interface{}) {
	// 将可变参数的值赋给空列表name, 可变参数本质是一个slice
	fmt.Println("可变参数args中的参数个数为:", len(args))
	name = args
	fmt.Println("遍历列表name:")
	for index, data := range name {
		if index == 5 {
			if value, ok := data.(rune); ok == true {
				fmt.Printf("name[%d] = %v\n", index, string(value))
			}
		} else {
			fmt.Printf("name[%d] = %v\n", index, data)
		}
	}
	fmt.Println("遍历可变参数args:")
	for index, data := range args {
		if index == 5 {
			if value, ok := data.(rune); ok == true {
				fmt.Printf("name[%d] = %v\n", index, string(value))
			}
		} else {
			fmt.Printf("name[%d] = %v\n", index, data)
		}
	}
}
