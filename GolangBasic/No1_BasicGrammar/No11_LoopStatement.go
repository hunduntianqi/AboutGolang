package main

import (
	"fmt"
	"time"
)

/*
	循环语句 ==> for 语句:
		格式:
			1. 无条件循环(无限循环)
				for {
					要循环执行的代码
				}
			2. 有条件循环(满足条件循环继续, 不满足条件循环中止)
				for 变量初始值; 循环条件; 变量值控制语句 {
					条件满足执行的代码
				}
				例:
				for i := 1; i < 10; i++ {
					fmt.Println(i)
				}
	控制语句使用的关键字:goto, break,continue
		goto:跳过其他代码块, 直接执行指定代码块,
			goto 指定代码块 在指定代码块之前, 代码块执行一次
			goto 指定代码块 在指定代码块之后, 代码循环执行多次(类似与无限循环)
		break:直接中止循环执行
		continue:跳过本轮循环, 进入下一次循环
*/

func main() {
	// for 语句使用
	var start = time.Now()
	fmt.Println(start)
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", i*j, "\t")
		}
		fmt.Println()
		// fmt.Println()
	}
	var end = time.Now()
	fmt.Println(end)
	time.Sleep(10 * time.Second)

	// goto用法
	//var i int = 1
	//One:
	//	fmt.Println("这是代码块1被执行, i的值为:", i)
	//	i++
	//	goto One
}
