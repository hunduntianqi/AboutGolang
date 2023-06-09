package main

/*
	网络协议:
		指网络中数据传输和数据解释的规则
		原始协议:
			指仅仅在一小部分人之间被遵守的协议
		标准协议:
			指原始协议被更多的人采用, 不断增加、改进、维护、完善, 最终形成的一个稳定, 完整的文件传输协议, 被广泛应用于
			各种文件传输中; 最早的ftp协议就是由此衍生的
	网络分层架构:
		为了减少协议设计的复杂性, 大多数网络采用分层的方式来组织, 每一层都有自己的功能, 每一层利用下一层提供的服务
		来为上一层提供服务, 本层服务的实现细节对上层屏蔽
		OSI/RM(理论上的标准):
			从上到下: 应用层-表示层-会话层-传输层-网络层-数据链路层-物理层
		TCP/IP(事实上的标准):
			从上到下: 应用层(对应理论层的: 应用层-表示层-会话层)-传输层-网络层-链路层(对应理论层的: 数据链路层-物理层)
		越下面的层, 越靠近硬件, 越上面的层, 越靠近用户
		各层功能:
			1. 链路层:
				通过网卡地址-MAC地址使得不同的计算机之间链接, 从而完成数据通信等功能, MAC地址:就是数据包的物理发送地址
				和物理接收地址
			2. 网络层:
				该层的作用是引进一套新的地址, 使得我们能够区分不同的计算机是否属于同一个子网络, 这套地址叫做网络地址,
				也是平时说的IP地址, 网络地址确定了计机所在的子网络, MAC地址将数据包传送到子网络中的目标网卡, 网络层协议包含的
				主要信息是源IP和目的IP
			3. 传输层:
				引入端口号, 相当于网络程序的一个编号, 包含源端口和目的端口, 确保由网络设备上的哪一个网络程序来发送和接收数据
				端口特点:
					3.1 对于同一个端口, 在不同系统中对应着不同进程
					3.2 对于同一个系统, 一个端口只能被一个进程拥有
			4. 应用层:
				规定应用程序的数据格式
	网络通信的条件:
		1. 网卡-MAC地址: 不需要用户处理
		2. 逻辑地址-IP地址: 需要用户指定, 是为了确定由哪个网络设备接收数据
		3. 端口: 为了确定由网络设备哪个网络程序来进行通信
*/

func main() {

}
