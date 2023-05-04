package main

import (
	"fmt"
	"reflect"
)

/*
	通过 reflect.TypeOf(variableName) 获取 reflect.Type对象
*/

func main() {
	typeI := reflect.TypeOf(1)
	typeS := reflect.TypeOf("hello")
	fmt.Println(typeI) //int
	fmt.Println(typeS) //string
}
