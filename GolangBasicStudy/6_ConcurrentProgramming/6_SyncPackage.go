package main

/*
   Sync包:
        1. sync.WaitGroup
            sync.WaitGroup内部维护着一个计数器, 计数器的值可以增加和减少
            方法:
            1.1 (wg *WaitGroup) Add(delta int): 将计数器的值 +delta
            1.2 (wg *WaitGroup) Done(): 将计数器的值 -1
            1.3 (wg *WaitGroup) Wait(): 阻塞程序运行, 直到计数器值 归0
            注意:
                sync.WaitGroup是一个结构体, 传递的时候要传递指针
        2. sync.Once:
            使用场景==>需要确保某些操作在高并发的场景下只执行一次
            只有一个Do方法: func (o *Once) Do(f func()) {}
            注意: 如果要执行的函数f需要传递参数就需要搭配闭包来使用
            定义sync.Once变量:
                var var_name sync.Once
            sync.Once内部包含一个互斥锁和一个布尔值:
                互斥锁: 保证初始化操作是并发安全的
                布尔值: 记录初始化是否完成
                这样设计可以保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次
        3. sync.Map:
            是Go语言中提供的一个并发安全版的map, 该map不需要像内置的Map一样使用make函数初始化即可使用
            定义sync.Map变量:
                var map_name = sync.Map{}
            sync.Map常用方法:
                Store(key, value):存储一对Key-value值
                Load(key): 根据key获取对应的value, 并判断该key是否存在
                LoadOrStore(key, value): 如果key对应的value存在, 则返回该value, 如果不存在, 存储相应的value
                Delete(key): 删除一个key-value键值对
                Range(f func(key, value) bool): 循环迭代sync.Map, 效果与for range一样
*/

import (
	"fmt"
	"strconv"
	"sync"
)

// 定义sync.WaitGroup变量
var wgDemo sync.WaitGroup

// 通过sync.WaitGroup实现并发同步
func syncDemo1() {
	wgDemo.Add(2) // 计数器 +2
	go func() {
		defer wgDemo.Done() // 每执行一次, 计数器 -1
		fmt.Println("Hello Goroutine One")
	}()

	go func() {
		defer wgDemo.Done()
		fmt.Println("Hello Goroutine Two")
	}()
	wgDemo.Wait() // 会阻塞直到计数器归0, 才会继续向下执行
	fmt.Println("main goroutine done!")
}

// 定义一个sync.Map变量
var sMap = sync.Map{}

// 使用sync.Map
func syncDemo2() {
	for i := 0; i < 20; i++ {
		wgDemo.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			sMap.Store(key, n)
			value, _ := sMap.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wgDemo.Done()
		}(i)
	}
	wgDemo.Wait()
}

func main() {
	syncDemo1()
	syncDemo2()
}
