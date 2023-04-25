package No8_ConcurrentProgramm

/*
	Timer:
		是一个定时器对象, 代表未来的一个单一事件
		创建Timer定时器:
			timer := time.NewTimer(指定时间)
			指定时间后会向timer通道(channel)写入当前时间
			time_now := <- timer.C ==> 定时器通道写入时间后, 取出数据
		通过Timer可以实现延时功能
		定时器停止:
			Timer.stop(): 定时器停止后, 不会再等待执行写入定时器当前时间
		定时器重置:
			Timer.Reset(指定时间): 修改原定时器指定时间, 原定时器指定时间失效
	Ticker:
		是一个定时触发的计时器, 会以一定会的时间间隔向channel中写入当前时间, 可以以固定的时间间隔从channel中读取数据
		创建定时器:
			ticker := time.NewTicker(指定间隔时间) // 定时器每隔一定时间向channel中写入一次数据(当前时间)
		停止定时器:
			Ticker.Stop(): 定时器停止后, 不会再循环向channel中写入数据
		重置定时器:
			Ticker.Reset(指定时间): 修改原定时器指定时间, 原定时器指定时间失效
*/
