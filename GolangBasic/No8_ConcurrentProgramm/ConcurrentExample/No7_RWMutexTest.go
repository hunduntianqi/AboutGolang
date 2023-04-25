package main

import (
	"fmt"
	"sync"
	"time"
)

// 同时定义多个全局变量
var (
	x         int
	waitGroup sync.WaitGroup
	rwLock    sync.RWMutex
)

// 定义函数, 加写锁操作共享资源
func write() {
	// 加写锁
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond)
	fmt.Println("写锁操作: x = ", x)
	// 解写锁
	rwLock.Unlock()
	waitGroup.Done()
}

// 定义函数, 加读锁操作共享资源
func read() {
	// 加读锁
	rwLock.RLock()
	x = x + 1
	time.Sleep(time.Millisecond)
	fmt.Println("读锁操作: x = ", x)
	// 解读锁
	rwLock.RUnlock()
	waitGroup.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go read()
	}

	waitGroup.Wait()
	end := time.Now()
	fmt.Println("共耗时", end.Sub(start))
}
