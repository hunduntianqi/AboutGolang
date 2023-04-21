package main

/*
	select语句:
		类似于switch语句, 具有一系列的case语句和一个默认的分支语句, 每个case会对应一个通道的发送或接收过程
		与switch语句相比, select语句有较多的限制, 其中最大的一条限制为: 每个case语句里必须是一个IO操作
		语法格式:
		select {
			case <- ch1:
				// ch1成功取出数据时, 要执行的语句
			case ch2 <- value:
				// 成功向ch2中存入数据时, 要执行的语句
			...
			default:
				// 以上case语句均不满足条件时执行的语句
		}
		select语句会一直等待, 直到某个case的通信操作成功完成时, 就会执行该case对应的语句, 如果一直没有case完成通信操作, select
		语句会一直等待, 如果存在default语句, 会执行default语句内容, 同时程序会从select语句后的语句中恢复
		1. select可以同时监听一个或多个channel, 直到其中一个channel ready(准备好)
		2. 如果select语句中有多个case语句同时满足条件, 则会随机选择一个case语句执行
		3. select语句可以判断channel是否存满
*/

import "fmt"

// 定义生产者函数, ch只写, quit只读
func fibonaqi(ch chan<- int, quit <-chan bool) {
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
			num := <-ch
			fmt.Println(num)
		}
		// 可以停止
		quit <- true
	}()
	// 生产者, 向channel中写入数据
	fibonaqi(ch, quit)
}
