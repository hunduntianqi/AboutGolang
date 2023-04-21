package main

/*
	Timer:
		是一个定时器对象, 代表未来的一个单一事件
		创建Timer定时器:
			timer := time.NewTimer(指定时间)
			指定时间后会向timer通道(channel)写入当前时间: time_now := <- timer.C, 定时器通道写入时间后, 取出数据不再阻塞
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

import (
	"fmt"
	"time"
)

func timerDemo() {
	// 创建一个定时器, 设置时间为2s, 2s后, 向timer通道写入当前时间
	timer := time.NewTimer(2 * time.Second)
	t_now := time.Now()
	fmt.Printf("当前时间为:%v\n", t_now)
	t := <-timer.C
	fmt.Printf("2s后: %s\n", t)
	fmt.Printf("t_Gap = %v\n", t.Sub(t_now))
	// <-timer.C // 此处会阻塞死锁
}

func tickerDemo() {
	// 创建定时器, 指定时间间隔为1s, 定时器每隔1s向channel中写入一次数据(当前时间)
	ticker := time.NewTicker(1 * time.Second)
	// ticker.Reset(2 * time.Second) // 重置定时器时间间隔为2s
	i := 0
	go func() {
		for {
			t := <-ticker.C
			fmt.Printf("i = %d\n", i)
			fmt.Println(time.Now(), t)
			i++
			if i == 5 {
				ticker.Stop() // 停止定时器
			}
		}
	}()

	for {
		if i == 5 {
			break
		}
	} // 死循环防止主协程停止
}

func main() {
	timerDemo()
	tickerDemo()
}
