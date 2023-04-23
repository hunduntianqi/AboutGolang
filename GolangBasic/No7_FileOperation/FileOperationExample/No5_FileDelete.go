package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
	递归删除文件夹
*/

func main() {
	// 定义变量, 储存文件路径
	var path string
	fmt.Println("请输入要删除文件夹路径: ")
	fmt.Scan(&path)
	// 调用函数, 删除文件
	fileDelete(path)
}

// 定义函数, 执行文件删除功能
func fileDelete(filePath string) {
	// 获取文件对象
	file, errFile := os.Open(filePath)
	if errFile != nil {
		fmt.Println(errFile)
		return
	}
	// 获取文件对象信息
	fileInfo, errInfo := file.Stat()
	if errInfo != nil {
		fmt.Println(errInfo)
		return
	}
	// 判断文件是否为文件夹
	if fileInfo.IsDir() {
		// 是文件夹, 获取文件夹所有文件路径信息
		fileAllPath, errAllPath := filepath.Glob(filePath + "//*")
		if errAllPath != nil {
			fmt.Println(errAllPath)
			return
		}
		// 遍历文件路径列表
		for _, fileChildPath := range fileAllPath {
			// 获取子文件对象
			childFile, errChildFile := os.Open(fileChildPath)
			if errChildFile != nil {
				fmt.Println(errChildFile)
				return
			}
			// 获取子文件对象信息
			childInfo, errChildInfo := childFile.Stat()
			if errChildInfo != nil {
				fmt.Println(errChildInfo)
				return
			}
			// 判断子文件对象是否为文件夹
			if childInfo.IsDir() {
				// 是文件夹, 递归调用
				fileDelete(fileChildPath)
			} else {
				// 不是文件夹删除文件
				childFile.Close()
				os.Remove(fileChildPath)
			}
		}
		// 子文件处理完毕, 删除文件对象
		file.Close()
		os.Remove(filePath)
	} else {
		// 不是文件夹, 直接删除文件
		file.Close()
		os.Remove(filePath)
	}
}
