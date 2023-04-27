package main

import (
	"fmt"
	"net"
)

/*
	测试UDP服务端代码
*/

func main() {
	fmt.Println("UDP服务端程序启动~~")
	// 创建UDPAddr类型变量
	udpAddr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}
	// 创建UDP服务端监听对象
	udpConn, errUDP := net.ListenUDP("udp", &udpAddr)
	if errUDP != nil {
		fmt.Println(errUDP)
		return
	}
	for {
		// 定义字节切片, 存储接收数据
		receiveData := make([]byte, 1024)
		// 读取数据
		n, addr, errRead := udpConn.ReadFromUDP(receiveData)
		if errRead != nil {
			fmt.Println("数据读取失败 ==> ", errRead)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(receiveData[:n]), addr, n)
	}
	// 关闭连接
	udpConn.Close()
}
