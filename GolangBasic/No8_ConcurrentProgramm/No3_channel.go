package No8_ConcurrentProgramm

/*
	管道 ==> channel
		Go语言提倡通过共享通信来共享内存,而不是通过共享内存实现通信
		channel 是 Golang 中的一种特殊引用类型, 像一个传送带或者队列, 遵循先入先出, channel 可以让一个 goroutine 发送特定值到
		另一个 goroutine, 实现不同 goroutine 之间的通信
		channel定义格式:
			var variableName chan dataType
				variableName ==> channel 类型变量名
				chan ==> 表明声明的变量是一个 channel
				dataType ==> 规定 channel 中的数据类型
		channel 变量属于引用类型, 默认值为 nil, 需要初始化后才可以使用, 初始化格式如下:
			chanName = make(chan dataType, [缓冲大小]) ==> 其中缓冲大小是可选项, 用于指定该管道中可存入数据的个数
		也可以使用 make() 直接初始化管道, 不需要事先声明
			chanName := make(chan dataType, [缓冲大小]) ==> 其中缓冲大小是可选项, 用于指定该管道中可存入数据的个数
		channel 操作:
			channel有发送(send), 接收(receive)和关闭(close)三种操作
			发送和接收都使用"<-"符号
			1. 发送 ==> send: 将一个数据存入 channel, 格式: chanName <- value
			2. 接收 ==> 从 channel 中取出一个值, 格式: variableName := <- chanName, 忽略结果: <- chanName ==> 将读取的数据丢弃
			3. 关闭 ==> 关闭 channel, 格式: close(chanName)
			注意:
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
				var chanName chan<- type: 单向channel, 可以写入type类型的数据
				var chanName <-chan type: 单向channel, 可以取出type类型的数据
				chan<-: 表示数据存入管道
				<-chan: 表示将数据从管道中取出
			双向channel转换为单向channel:
				c := make(chan type)
				var send chan<- type = c: 只能将数据存入管道
				var send <-chan type = c: 只能从管道中取出数据
				注意: 单向channel不能再转换为普通的双向channel
			单向channel应用:作为函数参数传递
*/
