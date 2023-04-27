package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

/*
	TCP 服务端代码测试
*/
// 定义 sync.WaitGroup 对象用于线程同步
var waitGoroutine sync.WaitGroup

func main() {
	// 创建 Listen 监听对象, 指定通信类型和绑定 IP
	listen, errListen := net.Listen(`tcp`, `localhost:8080`)
	if errListen != nil {
		fmt.Println("监听对象创建失败 ==> ", errListen)
		return
	}
	fmt.Println("TCP服务端程序启动~~~")
	for {
		// 监听客户端连接, 获取客户端连接对象
		clientConn, errClient := listen.Accept()
		if errClient != nil {
			fmt.Println("客户端连接失败 ==> ", errClient)
			return
		}
		// 子协程开启前, sync.WaitGroup 计数器加一
		waitGoroutine.Add(1)
		// 开启子协程, 处理客户端请求
		go func(conn net.Conn) {
			fmt.Println(conn.RemoteAddr().String(), "连接成功, 解析请求数据: ")
			// 延迟调用, 数据处理完毕后计数器减一
			defer waitGoroutine.Done()
			// 定义一个字节切片, 存储读取到的客户端数据
			clientData := make([]byte, 1024)
			for {
				num, err := conn.Read(clientData)
				if err == io.EOF {
					fmt.Println("客户端数据读取完毕")
					break
				} else {
					fmt.Println(conn.RemoteAddr().String(), ": ", string(clientData[:num]))
				}

				// 向客户端回复数据
				_, errReply := conn.Write([]byte("已收到数据, 目前连接正常!!"))
				if errReply != nil {
					fmt.Println("回复客户端数据发送失败 ==> ", errReply)
					return
				}
				// 处理完毕, 关闭连接
				defer func() {
					errClose := conn.Close()
					if errClose != nil {
						fmt.Println("客户端连接关闭异常 ==> ", errClose)
						return
					}
				}()
			}

		}(clientConn)
		// 阻塞主协程
		waitGoroutine.Wait()
		//if runtime.NumGoroutine() == 1 {
		//	fmt.Println("程序结束!")
		//	return
		//}
	}
}
