package main

import (
	"fmt"
	"sync"
)

// 定义全局变量
var num int = 0
var wg sync.WaitGroup
var numLock int = 0
var lock sync.Mutex

// 定义未加锁函数
func add() {
	for i := 0; i < 5000; i++ {
		num++
	}
	wg.Done()
}

// 定义加锁函数
func addLock() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		numLock++
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func main() {
	// 不加锁多goroutine操作共享资源
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println("num的最终结果是: ", num) // 若无数据竞争, num = 10000; 实际因为存在数据竞争, num的值无法确定
	// 加锁多goroutine操作共享资源
	wg.Add(2)
	go addLock()
	go addLock()
	wg.Wait()
	fmt.Println("numLock的最终结果是: ", numLock) // 加锁后同时只有一个goroutine操作公共数据, numLock的值为10000
}
