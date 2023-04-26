package main

import (
	"fmt"
	"net"
)

/*
	TCP 客户端代码测试
*/

func main() {
	// 创建客户端连接对象, 绑定服务端ip
	connClient, errClient := net.Dial(`tcp`, `localhost:8080`)
	if errClient != nil {
		fmt.Println("与服务端连接对象创建失败 ==> ", errClient)
		return
	}
	// 定义变量, 接收用户输入
	userInput := ""
	for {
		fmt.Println("请输入你要发送的信息: ")
		fmt.Scan(&userInput)
		// 向服务端发送数据
		_, errSend := connClient.Write([]byte(userInput))
		if errSend != nil {
			fmt.Println("向服务端发送数据失败 ==> ", errSend)
			return
		}
		if userInput == "quit" {
			// 关闭与服务端连接
			connClient.Close()
			return
		}
		// 定义切片, 接收服务端回复数据
		serverReply := make([]byte, 1024)
		// 读取数据
		connClient.Read(serverReply)
		fmt.Println("服务端回复数据 ==> ", string(serverReply))
	}
}
