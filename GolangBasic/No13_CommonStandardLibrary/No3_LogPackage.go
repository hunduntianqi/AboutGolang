package main

import (
	"fmt"
	"log"
	"os"
)

/*
	Log 包:
		golang 的 log 包定义了 Logger 类型,该类型提供了一些格式化输出的方法
			Print系列 ==> Print | Printf | Println
			Fatal系列 ==> Fatal | Fatalf | Fatalln
				会在写入日志信息后调用 os.Exit(1), 会使整个程序结束执行, 立即退出
			Panic系列 ==> Panic | Panicf | Panicln
				会在写入日志信息后panic, 报错
		配置logger:
			func Flags() int ==> 返回标准logger的输出配置
			func SetFlags(flag int) ==> 设置标准logger的输出配置, 可同时设置多个选项
				flag选项是log标准库提供的一系列定义好的常量
					Ldate ==> 输出日期信息
					Ltime ==> 输出时间信息
					Lmicroseconds ==> 输出微秒级别的时间信息
					Llongfile ==> 输出文件的全路径名 + 行号
					Lshortfile ==> 输出文件名 + 行号, 会覆盖掉Llongfile
					LUTC ==> 使用 UTC 时间输出
					LstdFlags ==> 等价于 Ldate | Ltime, 是标准 logger 的初始值
				例: log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
		配置日志前缀:
			可以对日志信息添加指定的前缀, 方便之后对日志信息进行检索和处理
			func Prefix() string ==> 以字符串形式获取标准logger的输出前缀
			func SetPrefix(prefix string) ==> 设置标准logger的输出前缀
		配置日志输出位置:
			func SetOutput(w io.Writer) ==> 用来设置标准logger的输出位置, 默认是标准错误输出
			例: 将日志信息输出到文件中 ==> log.SetOutput(FileObject)
		自定义创建logger对象:
			func New(out io.Writer, prefix string, flag int) *Logger:
				out ==> 日志输出位置
				prefix ==> 输出日志的前缀
				flag ==> 输出日志的配置信息
*/

func main() {
	log.Println("这是一条很普通的日志~~~~")
	//log.Fatalln("这是一条会触发 Fatal 的日志~~~")
	//log.Panicln("这是一条会触发 Panic 的日志~~~")
	// 配置logger
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.LstdFlags)
	log.Println("这是一条很普通的日志~~~~")
	// 设置标准logger的输出前缀
	log.SetPrefix("[info]")
	log.Println("这是一条很普通的日志~~~~")
	// 获取标准logger的输出前缀
	fmt.Println(log.Prefix())
	// 设置日志信息输出到文件中
	// 创建 文件对象
	file, errFile := os.Create("No13_CommonStandardLibrary/logfile.txt")
	if errFile != nil {
		fmt.Println("文件创建失败 ==> ", errFile)
		return
	}
	// 设置日志信息写入文件
	log.SetOutput(file)
	log.Println("这是一条写入文件的普通的日志~~~~")
	log.Println()
	file.Close()
}
