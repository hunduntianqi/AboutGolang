package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	测试runtime包常用函数
*/

// 创建一个 sync.WaitGroup 类型变量, 用于协程同步
var waitRuntime sync.WaitGroup

// 定义函数, 用于开启子协程
func childGoroutine(number int) {
	// 函数执行完毕后, 计数器减一
	defer waitRuntime.Done()
	if number == 19 {
		// 出让CPU执行权限
		runtime.Gosched()
		fmt.Println("子协程执行 ", number+1, "次")
	} else if number == 49 {
		// 终止当前协程
		fmt.Println("终止当前协程, 此时是第 ", number+1, "个协程在执行")
		runtime.Goexit()
		fmt.Println("看看我打印了没有, 如果我打印了说明程序并未终止")
	} else {
		fmt.Println("子协程执行 ", number+1, "次")
	}
}

func main() {
	// 查看当前电脑逻辑CPU个数 ==> runtime.NumCPU()
	fmt.Println("本机的逻辑处理器个数为: ", runtime.NumCPU())
	// 设置最大执行任务CPU数 ==> runtime.GOMAXPROCS
	numCPU := runtime.GOMAXPROCS(1)
	fmt.Println("之前系统最大执行任务CPU数为: ", numCPU)
	for i := 0; i < 100; i++ {
		waitRuntime.Add(1) // 计数器加一
		// 创建子协程
		go childGoroutine(i)
	}
	// 获取当前正在运行的协程数
	fmt.Println("当前系统正在运行的协程数为: ", runtime.NumGoroutine())
	// 阻塞主协程运行
	waitRuntime.Wait()
}
