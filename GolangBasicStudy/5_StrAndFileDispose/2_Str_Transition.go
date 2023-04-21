package main

import (
	"fmt"
	"strconv"
)

/*
   strconv包:
        实现了基本数据类型与字符串表示的转换
        常用函数:
            strconv.Atoi():
                用于将字符串类型的整数转化为int类型
                func Atoi(s string) (i int,err error)
                如果传入的字符串参数无法转化为int类型, 则会返回错误提示
            strconv.Itoa():
                用于将int类型数据转换为对应的string类型
                func Itoa(i int) string
            Parse系列函数:
                用于转换字符串为给定类型的值:
                ParseBool():
                    func ParseBool(str string) (value bool, err error)
                    返回字符串表示的bool值, 它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE; 否则返回错误
                ParseInt():
                    func ParseInt(s string, base int, bitSize int) (i int64, err error)
                    返回字符串表示的整数值, 接受正负号
                    base: 指定进制, 2~36; 如果base为0, 则会从字符串前置判断, "0x"是16进制, "0"是8进制, 否则是10进制
                    bitSize: 指定结果必须能无溢出赋值的整数类型; 0、8、16、32、64 分别代表 int、int8、int16、int32、int64
                ParseUnit():
                    func ParseUint(s string, base int, bitSize int) (n uint64, err error)
                    类似ParseInt但不接受正负号, 用于无符号整型
                ParseFloat():
                    func ParseFloat(s string, bitSize int) (f float64, err error)
                    解析一个表示浮点数的字符串并返回其值, 如果s合乎语法规则, 函数会返回最为接近s表示值的一个
                    浮点数(使用IEEE754规范舍入)
                    bitSize: 指定了期望的接收类型, 32是float32(返回值可以不改变精确值的赋值给float32), 64是float64
            Format系列函数:
                用于将其他基本类型转化为字符串类型
                FormatBool():
                    func FormatBool(b bool) string
                    根据b的值返回"true"或"false"
                FormatInt():
                    func FormatInt(i int64, base int) string
                    返回i的base进制的字符串表示, base 必须在2到36之间, 结果中会使用小写字母’a’到’z’表示大于10的数字
                FormatUint():
                    func FormatUint(i uint64, base int) string
                    是FormatInt的无符号整数版本
                FormatFloat():
                    func FormatFloat(f float64, fmt byte, prec, bitSize int) string
                    将浮点数表示为字符串并返回
                    bitSize: 32是float32, 64是float64
                    fmt: 表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、
                         ’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’
                         格式，否则’f’格式）
                    prec: 控制精度（排除指数部分）,对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如
                          果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f
            Append系列函数:
                将整数等转换为字符串后, 添加到现有的字节数组(切片)中
                AppendBool():
                    func AppendBool(dst []byte, b bool)
                    将“b”的值添加进dst字节切片中
                AppendInt():
                    func AppendInt(dst []byte, i int64, base int)
                    将 i 的 base 进制的整数添加进 dst 字节切片中
                AppendUint():
                    func AppendInt(dst []byte, i int64, base int)
                    AppendInt()的无符号整型版本
                AppendFloat()
                    func AppendFloat(dst[]byte, f float64, fmt byte, prec, bitSize int)
                    将浮点数数据添加进字节切片dst中
                    bitSize: 32是float32, 64是float64
                    fmt: 表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、
                         ’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’
                         格式，否则’f’格式）
                    prec: 控制精度（排除指数部分）,对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如
                          果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f
                AppendQuote()
                    func AppendQuote(dst []byte, s string)
                    将字符串 s 添加进字节切片dst中
*/

func AtoiDemo() {
	// 定义变量
	str := ""
	fmt.Println("请输入一个整数: ")
	fmt.Scan(&str)
	strInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("您输入的数字为: %d\n", strInt)
}

func ItoaDemo() {
	var num int
	fmt.Println("请输入一个整数: ")
	fmt.Scan(&num)
	numStr := strconv.Itoa(num)
	fmt.Printf("您输入的整数为: %s\n", numStr)
}

func parseSeriesDemo() {
	parseBool, errBool := strconv.ParseBool("1")
	if errBool != nil {
		fmt.Println(errBool)
	} else {
		fmt.Println("parseBool = ", parseBool)
	}

	parseInt, errInt := strconv.ParseInt("+123", 10, 0)
	if errInt != nil {
		fmt.Println(errInt)
	} else {
		fmt.Println("parseInt = ", parseInt)
	}

	parseUint, errUint := strconv.ParseUint("-123", 10, 0)
	if errUint != nil {
		fmt.Println(errUint)
	} else {
		fmt.Println("parseUint = ", parseUint)
	}

	parseFloat, errFloat := strconv.ParseFloat("1.452", 32)
	if errFloat != nil {
		fmt.Println(errFloat)
	} else {
		fmt.Printf("parseFloat = %0.2f\n", parseFloat)
	}
}

func formatSeriesDemo() {
	fmt.Printf("%s\n", strconv.FormatBool(true))
	fmt.Printf("%s\n", strconv.FormatInt(120, 10))
	fmt.Printf("%s\n", strconv.FormatInt(120, 10))
	fmt.Printf("%s\n", strconv.FormatFloat(120.5, 'f', 1, 32))
}

func appendDemo() {
	// 定义字节切片
	list := make([]byte, 0, 1024)
	list = strconv.AppendBool(list, false)
	// 方法中第二个参数为要追加的数据, 第三个参数为指定int进制类型, 下列为以10进制追加数据
	list = strconv.AppendInt(list, 100, 10)
	list = strconv.AppendQuote(list, "郭鹏涛")
	list = strconv.AppendFloat(list, 1.24, 'f', 2, 32)
	list = strconv.AppendQuote(list, "都去死吧！！")
	fmt.Printf("list = %v \n", string(list)) // 需要转换为字符串再打印,否则会以ascll编码数字显示
}

func main() {
	// AtoiDemo()
	// ItoaDemo()
	// parseSeriesDemo()
	// formatSeriesDemo()
	appendDemo()
}
