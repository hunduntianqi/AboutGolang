package main

import (
	"fmt"
	"time"
)

func timerDemo() {
	// 创建一个定时器, 设置时间为2s, 2s后, 向timer通道写入当前时间
	timer := time.NewTimer(2 * time.Second)
	tNow := time.Now()
	fmt.Printf("当前时间为:%v\n", tNow)
	t := <-timer.C
	fmt.Printf("2s后: %s\n", t)
	fmt.Printf("tGap = %v\n", t.Sub(tNow))
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
