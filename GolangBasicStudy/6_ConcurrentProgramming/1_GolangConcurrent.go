package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
   并发介绍:
        进程和线程:
            进程: 是程序在操作系统中的一次执行过程, 系统进行资源分配和调度的一个独立单位
            线程: 是进程的一个执行实体,是CPU调度和分派的基本单位,它是比进程更小的能独立运行的基本单位
            进程和线程的关系: 一个进程可以创建和撤销多个线程, 同一个进程中的多个线程之间可以并发执行
        并发和并行:
            并发: 多线程程序在一个核的cpu上运行, 主要由切换时间片来实现"同时运行"
            并行: 多线程程序在多个核的cpu上运行, 直接利用多核实现多线程的运行, Go语言中可以设置使用核数, 以发挥多核计算机的能力
        协程和线程:
            协程: 独立的栈空间, 共享堆空间, 调度由用户自己控制, 本质上有点类似于用户级线程, 这些用户级线程的调度也是自己实现的
            线程和协程的关系: 一个线程上可以跑多个协程, 协程是轻量级的线程
    Goroutine:
        由Go语言官方实现的超级"线程池"
        Goroutine奉行通过通信来共享内存, 而不是共享内存来通信
        goroutine的概念类似于线程, 是由Go的运行时(runtime) 调度和管理的; Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU
        Go语言中并发实现: 将相关代码包装成一个函数, 然后由goroutine去执行这个函数
        使用goroutine:
            Go语言中使用goroutine, 只需要在调用函数时在函数名前加上"go"关键字, 即可为函数创建一个goroutine(子协程)
            即: go func_name()
        注意:
            一旦主协程结束, 子协程也会被强制结束, 往往一般主协程是比较快的, 如果不做特殊处理, 在子协程来不及执行时, 主协程就已经
            执行完毕！！
            解决方法:
				sync.WaitGroup类型:
				使用方法:
					1. 定义WaitGroup变量
						var wg WaitGroup
					2. 调用Add()函数添加计数
					3. Done()函数减去一个计数
					3. 调用Wait()函数等待子协程执行完毕
    runtime包:
        1.runtime.Gosched():
            让出当前goroutine的执行权限, 调度器安排其他等待的任务执行, 并在下次某个时刻从该位置恢复执行
        2. runtime.Goexit():
            退出当前协程, 中止当前所在协程的运行
        3. runtime.GOMAXPROCS():
            设置执行代码的CPU核数,Go1.5版本之前, 默认使用的是单核心执行; Go1.5版本之后, 默认使用全部的CPU逻辑核心数
            runtime.GOMAXPROCS(num) ==> num是要是用CPU的核数
    Go语言中操作系统线程和goroutine的关系:
        1. 一个操作系统线程(OS线程)对应用户态多个goroutine
        2. Go语言程序可以同时使用多个操作系统线程(OS线程)
        3. goroutine和OS线程是多对多的关系, 即m:n
*/

// 定义函数
func hello() {
	fmt.Println("Hello World goroutine!!")
}

var wg1 sync.WaitGroup

func hello2(i int) {
	if i == 3 {
		wg1.Done()
		runtime.Goexit() // 中止子协程运行, 该协程不会再继续向下执行代码
		fmt.Println("Hello Goroutine + ", i)
	} else {
		defer wg1.Done() // 等待goroutine结束就-1
		fmt.Println("Hello Goroutine + ", i)
	}

}

func main() {
	runtime.GOMAXPROCS(4) // 设置同时执行任务的核数为4
	start := time.Now()
	go hello()
	for i := 0; i < 10; i++ {
		wg1.Add(1) // 启动一个goroutine就登记加1
		go hello2(i)
	}
	wg1.Wait()        // 等待所有登记的goroutine结束, 再向下执行代码
	runtime.Gosched() // 从此处开始暂停主协程, 等待其他子协程执行完毕, 再执行主协程
	fmt.Println("Hello World main goroutine")
	// time.Sleep(3 * time.Second)
	end := time.Now()
	fmt.Printf("开始时间是:%v, 结束的时间是%v\n", start, end)
}
