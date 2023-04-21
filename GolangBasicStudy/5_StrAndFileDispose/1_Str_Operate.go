package main

import (
	"fmt"
	"strings"
)

/*
	字符串操作:
		strings包提供了一些对于字符串的操作函数
		字符串常用操作函数:
			1. Contains:
				func Contains(s, substr, string) bool:判断字符串s中是否包含字符串substr, 返回bool值
			2. Join:
				func Join(a []string, sep string) string:将切片a中的字符串通过字符sep链接为一个新的
                    字符串并返回
			3. Index:
				func Index(s, sep string) int: 在字符串s中查找字符sep对应索引, 返回索引值, 如果没有返回-1
			4. Repeat:
				func Repeat(s string, count int) string:将字符串s复制count次后返回新的字符串
			5. Replace:
				func Replace(s, old, new string, n int) string: 在s字符串中, 用new字符串替换old字符串,
                    n表示替换次数, 小于0表示全部替换
			6. Split:
				func Split(s, sep string) []string:将字符串s通过sep字符进行切割, 生成一个字符串切片返回
			7. Trim:
				func Trim(s string, cutset string) string: 在字符串s中头部和尾部去除cutset指定的字符串
			8. Fields:
				func Fields(s string) []string: 去除字符串s中的空格符, 并按照空格分割返回字符串类型切片
*/

func main() {
	// 定义字符串
	str := "这是一个用来测试的字符串！！"
	// strings.Contanis
	fmt.Println(strings.Contains(str, "测试")) // 判断字符串中是否包含"测试"
	// 定义切片
	list := make([]string, 0, 1024)
	list = append(list, "混", "沌", "天", "炁")
	// strings.Join
	fmt.Println(strings.Join(list, "-")) // 将切片中的字符串以“-”连接后返回
	// strings.Index
	fmt.Println(strings.Index(str, "一")) // 查找字符串"一"对应的索引, 如果没有返回-1
	// strings.Repeat
	fmt.Println(strings.Repeat(strings.Join(list, "-"), 5))
	// strings.Replace
	fmt.Println(strings.Replace(str, "测试", "检查", 1))
	// strings.Split
	fmt.Println(strings.Split(strings.Join(list, "-"), "-"))
	// strings.Trim
	fmt.Println(strings.Trim(str, "！"))
	// strings.Fields
	fmt.Println(strings.Fields(strings.Replace(str, "测试", "  ", 1)))

}
