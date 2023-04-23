package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
	filepath 包:
		filepath包实现了兼容各操作系统的文件路径相关的实用操作函数
		常用函数:
			1. func IsAbs(path string) bool: 判断指定路径是否是一个绝对路径
			2. func Abs(path string) (string, error): 返回指定路径地址代表的绝对路径, 如果path不是绝对路径, 会加入当前工作目录
														使之成为绝对路径
			3. func Join(elem ...string) string: 将任意数量的路径元素以路径分隔符拼接为一个单一路径返回
			4. func FromSlash(path string) string: 将指定路径中的斜杠（'/'）替换为路径分隔符并返回替换结果
			5. func ToSlash(path string) string: 将指定路径的路径分隔符替换为斜杠（'/'）并返回替换结果
			6. func Glob(pattern string) (matches []string, err error): 返回所有符合匹配字符串pattern的文件名称
				例: 获取某个文件夹下所有文件的绝对路径 ==> filepath.Glob(dirPath + "//*")
*/

func main() {
	//pathJudge()
	//getAbs()
	//testJoin()
	testGlob()
}

// 定义函数, 判断指定路径是否为绝对路径 ==> func IsAbs(path string) bool
func pathJudge() {
	// 定义变量, 存储路径信息
	path := ""
	fmt.Println("请输入一个路径信息:")
	fmt.Scan(&path)
	// 判断路径
	mark := filepath.IsAbs(path)
	fmt.Println("给定的路径信息是否为绝对路径: ", mark)
}

// func Abs(path string) (string, error): 返回指定路径地址代表的绝对路径
func getAbs() {
	// 定义变量, 存储路径信息
	path := ""
	fmt.Println("请输入一个路径信息:")
	fmt.Scan(&path)
	// 返回绝对路径信息
	absPath, errPath := filepath.Abs(path)
	if errPath != nil {
		fmt.Println(errPath)
		return
	}
	fmt.Println("给定路径信息的绝对路径是: ", absPath)
}

// func Join(elem ...string) string: 将任意数量的路径元素以路径分隔符拼接为一个单一路径返回
func testJoin() {
	// 定义变量, 存储路径信息
	path := ""
	fmt.Println("请输入一个路径信息:")
	fmt.Scan(&path)
	// 打开文件对象
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, errFileInfo := file.Stat()
	if errFileInfo != nil {
		fmt.Println(errFileInfo)
		return
	}
	// 判断文件是否为文件夹
	if fileInfo.IsDir() {
		// 是文件夹, 获取所有文件名切片
		fileAllPath, _ := filepath.Glob(path + "\\*")
		fileNum := len(fileAllPath)
		fileNameSlice, errFileNameSlice := file.Readdirnames(fileNum)
		if errFileNameSlice != nil {
			fmt.Println(errFileNameSlice)
			return
		}
		// 遍历文件名切片
		for _, fileName := range fileNameSlice {
			// 调用Join函数, 拼接绝对路径
			fmt.Println(filepath.Join(path, fileName))
		}
	} else {
		fmt.Println(filepath.Abs(path))
	}
}

// func Glob(pattern string) (matches []string, err error): 返回所有符合匹配字符串pattern的文件名称
func testGlob() {
	// 定义变量, 存储路径信息
	path := ""
	fmt.Println("请输入一个路径信息:")
	fmt.Scan(&path)
	// 打开文件对象
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, errFileInfo := file.Stat()
	if errFileInfo != nil {
		fmt.Println(errFileInfo)
		return
	}
	// 判断文件是否为文件夹
	if fileInfo.IsDir() {
		// 是文件夹, 获取所有文件路径切片
		fileAllPath, _ := filepath.Glob(path + "\\*")
		// 遍历文件路径切片
		for _, filePath := range fileAllPath {
			// 打印文件绝对路径
			absPath, _ := filepath.Abs(filePath)
			fmt.Println(absPath)
		}
	} else {
		fmt.Println(filepath.Abs(path))
	}
}
