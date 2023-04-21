package main

import "fmt"

/*
    匿名字段:
		体现了面向对象的继承特性
        指结构体中只有字段类型, 没有字段名称的字段, 所有的内置类型和自定义类型都是可以作为匿名字段使用的
        例:
            定义Person结构体:
                type Person struct {
                    name string
                    age int
                }
            定义Student结构体:
                type Student struct {
                    Person // 此处Person就是一个匿名字段
                    id int
                }
            继承匿名字段结构体初始化:
                student := Student{Person{name字段的值, age字段的值}, id字段的值}
                注意: 继承匿名字段的结构体在初始化时, 不能写字段名, 只能按字段顺序进行初始化赋值
			同名字段:
				指结构体中出现与匿名字段中已有字段同名的字段
				同名字段操作原则:就近原则, 现在自己结构体中找到对应字段进行操作,
				找不到再去嵌套字段中找对应字段进行操作
				指定字段操作:
					给自己的同名字段赋值: instance.同名字段 = value
					给匿名字段中同名字段赋值: instance.匿名字段.同名字段 = value
			其他匿名字段:
				非结构体匿名字段:
					type Student struct {
						Person // 结构体匿名字段
						int // 基础类型匿名字段
					}
				结构体指针类型匿名字段:
					type Student struct {
						*Person // 结构体指针匿名字段
						int // 基础类型匿名字段
					}
					指针字段初始化:
						结构体变量.Person = &Person{匿名字段成员赋值}
       方法的继承:
           Go语言中结构体的匿名字段体现了Java中的继承思想, 继承这个匿名字段的结构体同时继承了该匿名字段的成员和方法
       方法的重写(同名方法):
           例:
               结构体Student继承了匿名字段Person, Student实现了一个方法, Person中也实现了一个方法, 两个方法名字相同
           叫做方法重写或同名方法
           重写方法隐式调用特点-就近原则:
                变量名.方法名()
                先调用被结构体实现的方法, 找不到再去调用继承的匿名字段的方法
           重写方法显示调用继承重写方法:
               变量名.匿名字段.方法名() ==>指定调用的方法是匿名字段的方法
*/

// 定义类型Person
type Person struct {
	name string
	age  int
	sex  rune
}

// 为类Person绑定方法
func (person Person) printMessage() {
	fmt.Printf("name = %s, age = %d, sex = %s\n", person.name, person.age, string(person.sex))
}

// 定义类型Student
type Student struct {
	Person // 定义匿名字段Person, 则Student继承Person类中的成员以及为类Person绑定的方法
	id     int
	name   string // 定义同名字段
}

// 为类Student绑定方法
func (student Student) printMessage() {
	fmt.Printf("id = %d, name = %s, name_Person = %s, age = %d, sex = %s\n", student.id, student.name,
		student.Person.name, student.age, string(student.sex))
}

func main() {
	// 定义Student类型变量
	student := Student{Person{"混沌天炁", 24, '男'}, 1, "混沌天帝"}
	student.Person.printMessage() // 显示调用Person绑定的方法
	student.printMessage()        // 隐式调用Student绑定的方法
}
