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
		定时器 ==> time.Tick(时间间隔), 可以设置一个定时器, 本质上是一个管道 channel
			例:
				定义一个定时器, 时间间隔为 1s ==> ticker := time.Tick(time.Second)
				for tickName :=range ticker {
					// 每隔一秒执行的任务代码
				}
		时间格式化:
			使用 time.Time.Format() 方法可进行时间格式化, 将时间字符串转换为想要的形式, Golang的时间格式化
			模板不是常见的 Y-m-d H:M:S, 而是Go语言的诞生时间:2006年1月2号15点04分
			格式化模板:
				24小时制 ==> time.Time.Format("2006-01-02 15:04:05.000 Mon Jan")
				12小时制 ==> time.Time.Format("2006-01-02 03:04:05.000 PM Mon Jan")
				年月日在前, 时分秒在后 ==> time.Time.Format("2006/01/02 15:04")
				时分秒在前, 年月日在后 ==> time.Time.Format("15:04 2006/01/02")
				只要年月日信息 ==> time.Time.Format("2006/01/02")
		解析时间字符串:
			1. 加载时区 ==> zoneName, err := time.LoadLocation("timeZone")
			2. 按照指定时区和格式化模板解析时间字符串:
				timeObj, err := time.ParseInLocation("timeFormatStr", "timeString", zoneName)
					timeFormatStr ==> 时间格式化模板
					timeString ==> 要解析的时间字符串
					zoneName ==> 指定时区对象
				注意: 时间格式化模板必须与要解析的时间字符串格式一致, 否则会解析失败
*/

// 定义函数, 获取当前时间信息
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

// 定义函数, 测试和时间戳相关方法
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

// 定义函数, 测试时间操作
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

// 定义函数, 测试时间格式化
func timeFormat() {
	// 获取当前时间对象
	timeNow := time.Now()
	// 24小时制时间格式化
	fmt.Println(timeNow.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制时间格式化
	fmt.Println(timeNow.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	// 年月日时分秒格式化
	fmt.Println(timeNow.Format("2006-01-02 15:04:05"))
	// 时分秒年月日格式化
	fmt.Println(timeNow.Format("15:04:05 2006-01-02"))
	// 年月日格式化
	fmt.Println(timeNow.Format("2006/01/02"))
	// 时分秒格式化
	fmt.Println(timeNow.Format("15:04:05"))
}

// 定义函数, 解析时间字符串
func timeStringParse() {
	// 加载时区, 获取时区对象
	zoneName, errZone := time.LoadLocation("Asia/Shanghai")
	if errZone != nil {
		fmt.Println("时区对象加载失败 ==> ", errZone)
		return
	}
	// 定义一个时间字符串
	//timeStr := "2021:02:25 11:45"
	timeStr := "1998:02:24 06:00"
	// 解析时间字符串
	timeParse, errParse := time.ParseInLocation("2006:01:02 15:04", timeStr, zoneName)
	if errParse != nil {
		fmt.Println("解析时间字符串失败 ==> ", errParse)
		return
	}
	fmt.Println("解析获得的时间对象 ==> ", timeParse)
}

func main() {
	//timeTimeObject()
	//timeStamp()
	//timeOperator()
	//timeFormat()
	timeStringParse()
}
