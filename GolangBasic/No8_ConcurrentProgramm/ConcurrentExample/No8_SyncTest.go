package main

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
