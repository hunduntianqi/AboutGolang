package main

import (
	"fmt"
	"time"
)

/*
   time包:
      time包提供了时间的显示和测量用的函数, 日历的计算采用的是公历
      时间类型:
          time.Time表示时间类型, 通过time.Now()可以获取当前时间对象, 然后获取时间对象的年月日时分秒等信息
      时间戳:
          时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数, 它也被称为Unix时间戳（UnixTimestamp）
          时间对象获取时间戳:
              time_stamp := time.Now().Unix()
          获取纳秒时间戳:
              time_stamp := time.Now().UnixNano()
          time.Unix()函数可以将时间戳转换为时间格式:
              now := time.Unix(stamp, 0)
      时间操作:
          Add:
              func (t Time) Add(d Duration) Time:返回时间 t 增加时间间隔 d 后的时间
          Sub:
              两个时间之间的差值
              func (t Time) Sub(u Time) Duration: 返回一个时间段 t-u
          Equal:
              判断两个时间是否相同, 会考虑时区的影响, 因此不同时区标准的时间也可以正确比较
              func (t Time) Equal(u Time) bool
          Before:
              判断某一时间点是否在给定的时间点之前, 是则返回true, 否则返回false
              func (t Time) Before(u Time) bool
          After:
              判断某一时间点是否在给定的时间点之后, 是则返回true, 否则返回false
              func (t Time) After(u Time) bool
*/

func timeTimeObject() {
	// 获取当前时间对象
	nowTime := time.Now()
	fmt.Println("当前时间为: ", nowTime)
	year := nowTime.Year()     // 获取年信息
	month := nowTime.Month()   // 获取月信息
	day := nowTime.Day()       // 获取日信息
	hour := nowTime.Hour()     // 获取当前小时信息
	min := nowTime.Minute()    // 获取当前分钟信息
	second := nowTime.Second() // 获取当前秒信息
	fmt.Printf("%v-%d-%d %d:%d:%d\n", year, month, day, hour, min, second)
}

func timeStamp() {
	// 获取当前时间
	now := time.Now()
	// 获取时间戳
	timeStamp := now.Unix()
	fmt.Printf("当前时间的时间戳是: %d\n", timeStamp)
	// 纳秒时间戳
	timeNanoStamp := now.UnixNano()
	fmt.Printf("当前时间的纳秒时间戳是: %d\n", timeNanoStamp)
	// 时间戳转时间类型
	timeObject := time.Unix(timeStamp, 0)
	fmt.Printf("当前时间是: %v\n", timeObject)
}

func timeOperator() {
	// 获取当前时间
	now := time.Now()
	timeAdd := now.Add(1 * time.Hour) // 当前时间增加一小时
	fmt.Printf("当前时间增加一小时 = %v\n", timeAdd)
	timeSub := timeAdd.Sub(now) // 一小时后时间减去当前时间
	fmt.Printf("一小时后时间减去当前时间是: %v\n", timeSub)
	fmt.Printf("now和timeAdd是否相同? %v\n", now.Equal(timeAdd))
	fmt.Println(now.Before(time.Unix(100000000000000, 0)))
	fmt.Println(now.After(time.Unix(1000000000, 0)))
}

func main() {
	timeTimeObject()
	timeStamp()
	timeOperator()
}
