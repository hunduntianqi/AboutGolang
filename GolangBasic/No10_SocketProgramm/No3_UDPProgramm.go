package No10_SocketProgramm

/*
	UDP编程:
		UDP协议:
			UDP(User Datagram Protocol) 即用户数据报协议, 是一种无连接的传输层协议, 不需要建立连接直接进行数据发送和接收, 属于不可靠,
			没有时序的通信, 但是实时性比较好, 通常用于视频直播相关领域
		UDP服务端:
			1. 绑定ip与端口创建监听对象
				创建UDP服务端通信对象
					UDPConn, err := net.ListenUDP(network, &UDPAddrTypeObject)
						network ==> "udp" / "udp4" / "udp6"
						UDPAddr类型{
							IP: IP地址(net.IPV4() / net.IPV6())
							Port: portNumber
							Zone:  string  // IPv6范围寻址域
						}
				UDP连接对象常用方法:
					func (c *UDPConn) LocalAddr() Addr: 获取本地网络地址对象
					func (c *UDPConn) RemoteAddr() Addr: 获取远端连接网络地址对象
					func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err): 接收远端数据到字节切片, 返回接收数据字节数
						和数据来源地址对象, 与错误描述
					func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error): 向指定网络地址发送数据
			2. 监听对象接收数据并处理客户端通信
				接收数据 ==> func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err)
				发送数据 ==> func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
			3. 关闭监听对象
				func (c *UDPConn) Close() error: 关闭连接
		UDP客户端:
			1. 绑定ip与端口创建UDP客户端通信对象
				UDPConn, err := net.DialUDP(network, &UDPAddrTypeObject)
					network ==> "udp" / "udp4" / "udp6"
					UDPAddr类型{
						IP: IP地址(net.IPV4() / net.IPV6())
						Port: portNumber
						Zone:  string  // IPv6范围寻址域
					}
			2. 向服务端发送数据与接收服务端回复
				接收数据 ==> func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err)
				发送数据 ==> func (c *UDPConn) Write(b []byte) (int, error): 发送数据到远端连接对象
			3. 关闭客户端连接对象
				UDPConn.Close()
*/
