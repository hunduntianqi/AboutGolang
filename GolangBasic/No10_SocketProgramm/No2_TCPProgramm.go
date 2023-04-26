package No10_SocketProgramm

/*
	TCP编程:
		TCP协议:
			TCP/IP(Transmission Control Protocol/Internet Protocol), 即传输控制协议 / 网间协议, 是一种面向连接的, 可靠的, 基于
			字节流的传输层通信协议, 因为是面向连接的协议, 数据像水流一样传输, 会存在黏包问题
		TCP 服务端:
			一个 TCP 服务端可以同时连接很多个客户端
			TCP服务端程序处理流程:
				1. 监听端口
					listen, err := net.Listen(network, address)
						network ==> 网络类型(tcp)
						address ==> IP地址和端口, 本地地址可省略不写
						listen对象常用方法:
							1. listen.Accept() ==> 获取一个客户端连接对象
							2. listen.Close() ==> 关闭监听器对象
				2. 接收客户端请求建立连接
					conn, err := listen.Accept() ==> 获取一个客户端连接对象
					客户端连接对象常用方法:
						a. conn.Read(b []byte) (n int, err error): 读取客户端发送数据到一个字节切片中, 返回读取的字节数和错误描述,
							当数据读取完毕后在读取数据后会返回 err = io.EOF, 可作为数据读取结束标志
						b. conn.LocalAddr().Network(): 获取当前连接的网络类型(tcp /udp)
						c. conn.LocalAddr().String(): 获取当前服务端绑定地址与端口信息的字符串形式数据
						d. conn.RemoteAddr().String(): 获取远程连接客户端的 IP地址与端口信息 的字符串形式数据
						e. conn.Close(): 关闭客户端连接对象
						f. conn.Write(b []byte) (n int, err error): 向客户端连接对象发送数据, 返回发送成功数据的字节数和错误描述
				3. 创建子协程 (goroutine) 处理客户端请求
					创建函数或匿名函数, 编写客户端请求数据处理代码
		TCP 客户端:
			客户端通信流程:
				1. 与服务端建立连接
					connClient, err := Dial(network, address string):
						network ==> 指定通信协议, tcp / udp
						address ==> 要建立连接的 ip地址与端口, 一般指服务端ip与端口
				2. 与服务端进行数据传输(通信):
					向服务端发送数据 ==> connClient.Write(b []byte) (n int, err error)
					接收服务端回复数据 ==> connClient.Read(b []byte) (n int, err error)
				3. 关闭连接
					connClient.Close()
*/
