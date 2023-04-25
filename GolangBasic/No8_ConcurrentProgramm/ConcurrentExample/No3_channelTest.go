package main

import "fmt"

// 定义函数作为接收者
func chanReceiver(ch chan int) {
	receive := <-ch
	fmt.Printf("接收成功==>receive = %d\n", receive)
}

// 定义函数作为发送者
func chanSend() {
	// 定义一个无缓冲channel
	ch := make(chan int)
	go chanReceiver(ch) // 开启一个goroutine, 并调用接收者函数, 如果没有接收者, 在向无缓冲通道存入数据时, 会造成死锁异常(panic)
	ch <- 10            // 发送者向管道存入数据
	fmt.Println("发送成功 ch = ", ch)
}

// 定义函数操作有缓冲通道
func chanBuffer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	// 定义一个可存入5个int数据的缓冲通道
	ch := make(chan int, 5)
	// 向channel中发送数据
	ch <- 0
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Printf("通道ch中的元素个数为:%d, 容量为:%d\n", len(ch), cap(ch))
	// ch <- 5 // 达到channel缓冲大小后, 再向channel中存入数据会造成阻塞
	// 从通道中取出数据
	for i := 0; i < 5; i++ {
		fmt.Printf("chan[%d] = %d\n", i, <-ch)
	}
	fmt.Printf("通道ch中的元素个数为:%d, 容量为:%d\n", len(ch), cap(ch))
	// 通道为空, 再取出数据会造成死锁阻塞
	// <-ch // fatal error: all goroutines are asleep - deadlock!
	// 对一个关闭的并且没有数据的通道执行接收操作会得到对应类型的零值
	close(ch)
	fmt.Printf("通道关闭且通道为空, 取出数据为类型零值==> ch = %d\n", <-ch)
	fmt.Printf("通道ch中的元素个数为:%d, 容量为:%d\n", len(ch), cap(ch))
	// 对一个关闭的通道再存入数据会引发panic异常
	ch <- 10 // send on closed channel
	// 对已经关闭的channel再执行关闭操作, 会引发panic异常
	close(ch) // close of closed channel
}

// 定义单向管道只可以存入数据
func receiveSignChan(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("管道中存入数据: ", i)
	}
	close(ch)
}

// 定义函数, 参数为仅可以取出数据的单向管道
func sendSignChan(chSend <-chan int) {
	// 便利管道
	for i := 0; i < 10; i++ {
		fmt.Println("从管道中读取到数据: ", <-chSend)
	}
}
func main() {
	chanSend()
	chanBuffer()
	// 定义一个管道
	var ch chan int = make(chan int, 10)
	// 调用函数, 当做仅能发送数据管道传入参数, 向管道中写入数据
	receiveSignChan(ch)
	// 调用函数, 当做仅能接收数据管道传入参数, 从管道中读取数据
	sendSignChan(ch)
}
