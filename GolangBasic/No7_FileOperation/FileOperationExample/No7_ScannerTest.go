package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建只读文件对象
	file, errFile := os.Open("No7_FileOperation/FileOperationExample/Test.txt")
	if errFile != nil {
		fmt.Println(errFile)
		return
	}
	// 包装文件对象创建Scanner对象
	scanner := bufio.NewScanner(file)
	// 读取文件数据
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println(string(scanner.Bytes()))
	}
	file.Close()
}
