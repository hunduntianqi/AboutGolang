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
