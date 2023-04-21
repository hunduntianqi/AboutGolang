package main

import (
	"fmt"
)

/*
	条件语句-if, if...else,嵌套if, if...else
		格式:
			if 判断条件表达式 {
				条件成立要执行的代码
			} else {
				条件不成立执行的代码
			}
	选择语句:switch,select
		switch语句:
			1. switch语句用于基于不同条件执行不同的动作, 每一个case分支都是唯一的, 从上至下逐一测试, 直到匹配为止
			2. switch语句执行过程从上到下
			3. switch语句还可以被用于type-switch来判断某个interface变量中存储的变量类型
			格式:
				switch 条件语句 {
					case 匹配值1:
						成功匹配后执行的代码
					case 匹配值2:
						成功匹配后执行的代码
					...
					default:
						所有匹配值都不满足条件后执行的代码
				}
				注意:Go语言中每个case最后都默认带有break,要想继续执行下面的代码, 可以在case末尾使用fallthrough关键字
				fallthrough:
					使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true
				Go语言中switch语句支持多条件匹配:
					switch 条件语句 {
							case 条件值1, 条件值2, 条件值3,...
							default
						}
		select语句:
			是Go语言中的一个控制结构, 类似于用于通信的switch语句, 每个case必须是一个通信(IO)操作, 要么是发送, 要么是接收
			语法:
				select {
					case <- chan1:
						// chan1成功读取数据后执行的内容
					case chan2 <- 1:
						// 成功向chan2写入数据后执行的内容
					default:
						// 以上条件均不满足后执行的内容
				}
			在一个select语句中, Go语言会按顺序从头至尾评估每一个发送和接收的语句, 如果其中的任意一条语句可以继续执行(没有被阻塞),
			就从这些可以执行的语句中任意选择一条来使用
			如果没有任意一条语句可以执行(所有case语句均被阻塞), 存在两种可能:
				1. 有default语句: 会执行default语句内容, 同时程序会从select语句后的语句中恢复
				2. 没有default语句: select语句将被阻塞, 直到至少有一个case语句可以进行下去
*/

func main() {
	var number int
	fmt.Println("请输入一个数字:")
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}
	// if 条件语句使用
	if number > 3 {
		fmt.Println("a的值是大于三的")
	} else {
		fmt.Println("a的值不大于三")
	}

	// switch选择语句使用
	switch 5 / 2 {
	case 1:
		fmt.Println("计算值为1")
	case 2:
		fmt.Println("计算值为2")
	case 3:
		fmt.Println("计算值为3")
	case 4:
		fmt.Println("计算值为4")
	case 5:
		fmt.Println("计算值为5")
	default:
		fmt.Println("所有计算结果都不对！！")
	}
	// switch语句判断变量类型
	var a interface{} = "3.7" // 空接口类型可以存储任意类型的数据
	switch a.(type) {
	case int:
		fmt.Println("变量的类型为整型")
	case float32:
		fmt.Println("变量的类型为浮点型")
	case string:
		fmt.Println("变量的类型为字符串类型")
	default:
		fmt.Println("以上类型都不对！！")
	}
}
