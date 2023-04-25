package main

import "fmt"

// 定义生产者函数, ch只写, quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		// 监听channel数据流动
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag =", flag)
			return
		}
	}
}
func main() {
	/*
		使用select实现斐波那契数列:
			即从第三个数字开始, 每一个数字都等于前两个数字相加
			1 1 2 3 5 8 13 21 ...
	*/
	// 定义channel保存数据
	ch := make(chan int)    // 进行数据通信
	quit := make(chan bool) // 判断程序是否结束
	// 消费者, 从channel中读取数据
	// 新建协程
	go func() {
		for i := 0; i < 10; i++ {
			num := <-ch // 从管道中取出数据
			fmt.Println(num)
		}
		// 可以停止
		quit <- true
	}()
	// 生产者, 向channel中写入数据
	fibonacci(ch, quit)
}
