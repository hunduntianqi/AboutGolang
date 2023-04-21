package main

import "fmt"

/*
	channel类型:
		Go 语言中的通道(channel)是一种特殊的类型, 像一个传送带或者队列, 遵循先入先出原则, 保证收发数据的顺序
		channel是一种引用类型, 类型零值为nil, 声明channel格式如下:
			var chan_name chan type
			chan_name: channel类型变量名
			chan: 定义该变量是一个channel
			type: 定义channel中存储的数据类型
			声明后的channel需要使用make关键字初始化后才可以使用:
				chan_name = make(chan type, [capacity]) ==> capacity是可选参数, 表示该channel中可存入数据的个数
			也可以使用make关键字创建channel, 直接初始化
			chan_name := make(chan type, [capacity])
		channel操作:
			channel有发送(send), 接收(receive)和关闭(close)三种操作
			发送和接收都使用"<-"符号
			发送(send): chan <- value ==> 将一个值存入通道
			接收(receive): <- chan 或 receiver = <- chan ==> 从通道中取出数据丢弃或赋值给其他变量
			关闭(close): close(chan) ==> 即关闭通道
				通道被关闭后:
					1. 对一个关闭的通道再发送数据就会导致panic
					2. 对一个关闭的通道进行接收会一直获取数据直到通道为空
					3. 对一个关闭的并且没有数据的通道执行接收操作会得到对应类型的零值
					4. 关闭一个已经关闭的通道会导致panic
		无缓冲通道:
			又被称为阻塞的通道, 只有接收者存在时, 才能发送数据到通道, 如果没有接收者, 在向channel发送数据时会造成死锁异常, 程序崩溃
			使用无缓冲通道进行通信将导致发送者和接收者的goroutine同步化, 因此, 无缓冲通道也被称为同步通道
		有缓冲通道:
			在定义channel时, 可以设置capacity大于0, 缓冲数据大于0的channel就是缓冲通道
			阻塞存入数据: 当缓冲通道中的数据个数达到设置的capacity时, 再向channel中存入数据会造成channel阻塞
			取出数据阻塞: 当缓冲通道为空时, 再从channel中取出数据, 会造成channel阻塞
			对于有缓冲通道:
				len()函数可以获取通道内元素的数量
				cap()函数可以获取通道的容量
		使用range遍历管道内容:
			遍历结束或会自动关闭循环
			for num := range ch{
			}
			value, ok := <- ch ==> 也可判断管道是否关闭
				如果管道关闭, 则ok = false
				管道未关闭, 则ok = true
		单方向channel:
			默认情况下, channel是双向的, 既可以向channel中写入数据, 也可以从channel中取出数据
			单向channel变量声明:
				var ch1 chan<- type: 单向channel, 可以写入type类型的数据
				var ch1 <-chan type: 单向channel, 可以取出type类型的数据
				chan<-: 表示数据存入管道
				<-chan: 表示将数据从管道中取出
			双向channel转换为单向channel:
				c := make(chan type)
				var send chan<- type = c: 只能将数据存入管道
				var send <-chan type = c: 只能从管道中取出数据
				注意: 单向channel不能再转换为普通的双向channel
			单向channel应用:作为函数参数传递
*/

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
func receiveSignChan() {
	var ch chan<- int
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

// 定义函数, 参数为尽可以取出数据的单向管道
func sendSignChan(ch_send <-chan int) {
	// 便利管道
	for ch_value := range ch_send {
		fmt.Println(ch_value)
	}
}
func main() {
	chanSend()
	chanBuffer()
}
