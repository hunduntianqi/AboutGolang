package main

import "fmt"

/*
   接口(interface):
       定义了一个对象的行为规范, 只定义规范, 不实现, 由具体的对象来实现规范的细节, 接口体现了面向对象的多态特征
        接口类型:
            在Go语言中, 接口(interface)是一种抽象的类型, interface是一组方法的集合, 接口做的事情就像是定义一个协议(规则),
            不关系属性, 只关心行为(方法)
        接口定义:
            type 接口类型名 interface {
                method_name1(param_list) (return_list) ()
                method_name2(param_list) (return_list) ()
                ...
            }
            接口名: 使用type将接口定义为自定义的类型名; Go语言的接口在命名时, 一般会在单词后面添加er, 如有写操作的接口叫Writer,
                   有字符串功能的接口叫Stringer等; 接口名最好要能突出该接口的类型含义
            method_name: 方法名, 当方法名和这个接口名首字母均是大写时, 这个方法可以被接口所在的包(package)之外的代码访问
            注意:
				1. 接口命名习惯以er结尾
				2. 接口只有方法声明, 没有实现, 没有数据字段
				3. 接口可以匿名嵌入其他接口, 或嵌入到结构中
        接口实现:
			接口是用来定义行为的类型, 这些被定义的行为不直接由接口实现, 而是由用户定义的类型实现, 一个实现了接口所有方法的具体类型就是
		    这个接口类型的实例, 换句话说, 接口就是一个需要实现的方法列表
            注意:
                1. 一个类型只有实现了接口中定义的全部方法, 才算实现了接口
			    2. 实现了接口方法的类型变量可以给接口类型变量赋值
            值接收者和指针类型接收者实现接口区别:
                1. 值接收者:
                    对于值接收者实现接口, 实现接口的 T 类型变量和 *T 类型变量均可以赋值给接口类型变量
                2. 指针类型接收者:
					对于指针类型接收者实现接口, 只有 *T 类型变量可以赋值给接口类型变量
		类型与接口的关系:
			1. 一个类型可以实现多个接口:
				一个类型可以同时实现多个接口, 而接口间彼此独立, 不知道对方的实现
			2. 一个接口可以有多个类型实现:
				不同的类型可以实现同一个接口
        接口类型变量:
            接口类型变量可以存储所有实现了该接口的实例
        接口嵌套:
			如果一个interface1作为interface2的一个嵌入字段, 那么interface2隐式的包含了interface1里面的方法
			例:
				type interface1 interface {
					sayHi()
					run()
				}
				type interface2 interface {
					interface1 // 相当于定义了sayHi()方法和run()方法
					study()
				}
		空接口:
			空接口不包含任何方法, 正因为如此, 所有的类型都实现了空接口, 因此空接口可以存储任意类型的数据
			当函数可以接受任意的对象实例时, 我们会将其声明为interface{}, 最典型的例子是标准库fmt中的PrintXXX系列的函数
			空接口应用:
				1. 作为函数的参数 ==> 使用空接口实现可以接收任意类型的函数参数
				2. 作为map的值 ==> 使用空接口实现可以保存任意值的字典
				3. 类型断言:
					接口值: 一个接口的值(简称接口值)是由一个具体类型和具体类型的值两部分组成的;
						   这两部分分别称为接口的动态类型和动态值
					类型断言语法格式:
						interface_var.(type)
						interface_var: 表示类型为interface{}的变量
						type: 表示断言变量可能是的类型
						该语法返回两个参数, 一个是该变量转换为要判定数据类型后的值, 一个是判断断言是否成功的布尔值, 若为true,
						则表示断言成功, false则表示断言失败; 该语法结合if语句或switch语句可以实现空接口的类型断言
		接口注意事项:
			只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口, 不要为了接口而写接口,
			那样只会增加不必要的抽象, 导致不必要的运行时损耗
*/

// 定义Sayer接口
type Sayer interface {
	say() // 定义方法say
}

// 定义类型Dog, Cat
type Dog struct{}
type Cat struct{}

// 类型Dog和Cat实现接口Sayer
func (d Dog) say() {
	fmt.Println("汪汪汪！！")
}

func (c Cat) say() {
	fmt.Println("喵喵喵！！")
}

// 定义测试函数 ==> 接口实现测试 与 实现了接口方法的类型变量可以给接口类型变量赋值测试
func interfaceTest1() {
	fmt.Println("=======接口实现测试 与 实现了接口方法的类型变量可以给接口类型变量赋值测试=======")
	// 定义Cat和Dog类型变量
	var c Cat
	var d Dog
	c.say()
	d.say()
	// 定义Sayer类型变量
	var s Sayer = c // 实现接口的类型变量可以给接口类型变量赋值
	s.say()
	s = d // 接口类型变量可以存储所有实现了该接口的实例
	d.say()
}

// 定义接口Move
type Move interface {
	move()
}

// 值接收者实现接口
func (d Dog) move() {
	fmt.Println("狗会移动")
}

// 指针类型接收者实现接口
func (c *Cat) move() {
	fmt.Println("猫可以上树")
}

// 定义测试函数 ==> 值接收者和指针类型接收者实现接口区别
func interfaceTest2() {
	fmt.Println("=======值接收者和指针类型接收者实现接口区别=======")
	// 定义Dog类型变量
	var dog Dog
	dog.move() //
	// 定义*Dog类型变量
	var dogPoint *Dog
	dogPoint = &dog
	dogPoint.move()
	// 定义Move接口变量
	var move Move
	move = dog      // Dog类型变量可以赋值给Move接口变量
	move = dogPoint // *Dog类型变量也可以赋值个Move接口类型变量
	move.move()
	// 定义Cat类型变量
	var cat Cat
	cat.move()
	// 定义*Cat类型变量
	var catPoint *Cat
	catPoint = &cat
	catPoint.move()
	// move = cat // Cat类型变量赋值给接口类型变量Move, 由于实现方式为指针接收者类型, 会报错
	move = catPoint
	move.move()
}

func main() {
	interfaceTest1()
	interfaceTest2()
}
