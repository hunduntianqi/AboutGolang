/*
		正则表达式:
			是一种进行模式匹配和文本操作的复杂而又强大的工具, 按照正则表达式的语法规则, 随需构造出的匹配规则能够从原始文本中
		筛选出几乎任何想要得到的字符组合
	        正则表达式匹配规则:
	            元字符:具有固定含义的特殊符号
	             常用元字符:
	                1.‘.’-匹配除换行符以外的任意字符
	                2.‘/w’-匹配字母或数字或下划线
	                3.‘/s’-匹配任意的空白符
	                4.‘/d’-匹配数字
	                5.‘/n’-匹配一个换行符
	                6.‘/t’-匹配一个制表符
	                7.‘^’-匹配字符串的开始
	                8.‘$’-匹配字符串的结尾
	                9.‘/W’-匹配非字母或数字或下划线
	                10.‘/D’-匹配非数字
	                11.‘/S’-匹配非空白符
	                12.‘a丨b’-匹配字符a或字符b
	                13.‘()’-匹配括号内的表达式,也表示一个组
	                14.‘[...]’-匹配字符组中的字符
	                15.‘[^...]’-匹配除了字符组中字符的所有字符
	             量词:控制前面的元字符出现的次数
	                1.‘*’-重复零次或更多次
	                2.‘+’-重复一次或更多次
	                3.‘？’-重复零次或一次
	                4.‘{n}’-重复n次
	                5.‘{n,}’-重复n次或更多次
	                6.‘{n,m}’-重复n到m次
	             贪婪匹配和惰性匹配:
	                1.“.*”/'.+'/'.?'-贪婪匹配:表示在整个表达式匹配成功的前提下, 尽可能多的匹配
	                2.“.*？”/'.+?'/'.??'-惰性匹配:表示在整个表达式匹配成功的前提下, 尽可能少的匹配
		Go语言通过regexp标准包为正则表达式提供了官方支持
		regexp包使用流程:
			1. 指定正则表达式匹配规则
				reg := regexp.MustComlile('正则匹配规则')
			2. 根据匹配规则提取信息
				reg.FindAllString(buf, -1):返回数据为只有一个参数的切片
				reg.FindAllStringSubmatch(buf, -1): 返回数据为多个切片组成的切片
*/
package main

import (
	"fmt"
	"regexp"
)

func main() {
	var buf string = "abc azc a7c aac 888 a9c tac"
	// 指定匹配规则
	reg, err := regexp.Compile(`a.c`)
	if err == nil {
		fmt.Println(reg)
	}
	fmt.Println(reg.FindAllString(buf, -1))
	fmt.Println(reg.FindAllStringSubmatch(buf, -1))
	for index, data := range reg.FindAllStringSubmatch(buf, -1) {
		fmt.Printf("data[%d]=%s\n", index, data[0])
	}
}
