package main

import (
	"fmt"
	"sync"
)

/*
	goroutine 使用测试
*/

// 定义 WaitGroup 类型全局变量
var waitTest sync.WaitGroup

// 定义测试函数
func goroutineTest(num int) {
	// 延迟调用sync.WaitGroup.Done(), 计数器减一
	defer waitTest.Done()
	fmt.Printf("子协程执行第 %d 次\n", num+1)
}

func main() {
	// for 循环创建多个子协程
	for i := 0; i < 100; i++ {
		// 调用方法 sync.WaitGroup.Add(1), 计数器加一
		waitTest.Add(1)
		// 使用 go 关键字, 开启子协程
		go goroutineTest(i)
	}
	// 调用 sync.WaitGroup.Wait() 方法, 阻塞主程序执行, 等待子协程执行完毕
	waitTest.Wait()
	fmt.Println("主协程执行完毕!!")
}
