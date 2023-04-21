package main

/*
   并发安全:
        在Go语言中, 可能会存在多个goroutine同时操作一个资源, 这种情况会发生资源竞争问题, 导致可能与预期不符
    互斥锁:
        是一种常用的控制共享资源访问的方法, 可以保证同一时间只有一个goroutine可以访问共享资源
        Go语言使用sync包的 Mutex 类 来实现互斥锁
            格式:
                var lock sync.Mutex // 定义一个锁
                lock.Lock() // 加锁
                公共代码
                lock.Unlock() // 解锁
                当有goroutine在操作被上锁的公共资源时, 其他Goroutine只有等待解锁后, 才可以操作公共资源,
                当有多个Goroutine等待一个锁时, 解锁后会随机唤醒一个goroutine操作公共代码
    读写互斥锁:
        读写锁:
            读锁: 当一个goroutine获取读锁时, 其他goroutine如果获取读锁, 会继续获得锁; 如果获取写锁, 则需要等待读锁释放
            写锁: 当一个goroutine获取写锁之后, 其他的goroutine无论是获取读锁还是写锁都会等待
        读写锁定义:
            Go语言使用sync包的 RWMutex 类实现
            var rwlock sync.RWMutex // 定义一个读写锁
            rwlock.Lock() // 加写锁
            // 需要上锁的代码
            rwlock.Unlock() // 解写锁
            rwlock.RLock() // 加读锁
            // 需要上锁代码
            rwlock.RUnlock() //解读锁
        注意:
            读写锁适合读多写少的场景, 如果读写操作差别不大, 就无法发挥读写锁的优势
*/

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
