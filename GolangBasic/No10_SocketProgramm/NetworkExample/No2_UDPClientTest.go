package main

import (
	"fmt"
	"net"
)

/*
	测试UDP服务端代码
*/

func main() {
	// 创建UDPAddr类型变量
	udpAddr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}
	// 创建UDP服务端监听对象
	udpConn, errUDP := net.DialUDP("udp", nil, &udpAddr)
	if errUDP != nil {
		fmt.Println(errUDP)
		return
	}
	// 向指定网络地址发送数据
	udpConn.Write([]byte("你好, 这是一条测试语句"))
	// 关闭连接
	udpConn.Close()
}
