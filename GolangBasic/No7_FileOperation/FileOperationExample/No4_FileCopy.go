package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
	文件复制, 考虑子文件夹, 递归复制
*/

func main() {
	// 定义变量, 记录源文件夹路径
	var sourceDirPath string
	fmt.Println("请输入源文件夹路径:")
	fmt.Scan(&sourceDirPath)
	// 定义变量, 记录目标文件夹路径
	var targetDirPath string
	fmt.Println("请输入目标文件夹路径:")
	fmt.Scan(&targetDirPath)
	// 调用函数, 获取源文件夹中所有文件路径信息
	sourceDirFilePath := getSourceDirFilePath(sourceDirPath)
	// 调用函数, 复制文件
	copyFile(sourceDirFilePath, sourceDirPath, targetDirPath)
}

// 定义函数, 获取源文件夹中所有文件路径信息
func getSourceDirFilePath(sourcePath string) []string {
	// 定义切片, 存储源文件对象路径
	sourceFileSlice := make([]string, 0)
	// 1. 判断源文件夹是否存在
	_, errSourceDir := os.Open(sourcePath)
	if errSourceDir == nil {
		// 源文件夹存在, 获取源文件夹下所有文件路径
		fileAllPath, errAllPath := filepath.Glob(sourcePath + "//*")
		if errAllPath != nil {
			fmt.Println(errAllPath)
			return nil
		}
		// 遍历切片
		for _, filePath := range fileAllPath {
			// 获取文件信息
			file, errFile := os.Open(filePath)
			if errFile != nil {
				fmt.Println(errFile)
				return nil
			}
			fileInfo, errFileInfo := file.Stat()
			if errFileInfo != nil {
				fmt.Println(errFileInfo)
				return nil
			}
			// 将文件路径加入切片
			sourceFileSlice = append(sourceFileSlice, filePath)
			fmt.Println("文件 ", filePath, " 已加入切片")
			// 判断对应文件是否为文件夹
			if fileInfo.IsDir() {
				// 是文件夹, 递归调用函数, 获取子文件夹对象路径
				childDirFilePath := getSourceDirFilePath(filePath)
				// 将子文件夹中的文件路径信息加入切片
				sourceFileSlice = append(sourceFileSlice, childDirFilePath...)
			}
		}
	} else {
		// 源文件夹不存在, 提示并结束程序
		fmt.Println(errSourceDir)
		return nil
	}
	return sourceFileSlice
}

// 定义函数, 复制源文件夹文件到目标文件夹
func copyFile(sourcePathSlice []string, sourceDirPath, targetDirPath string) {
	// 遍历切片
	for _, sourceFilePath := range sourcePathSlice {
		fmt.Println("文件路径信息为: ", sourceFilePath)
		// 获取源文件对象
		sourceFile, errFile := os.Open(sourceFilePath)
		if errFile != nil {
			fmt.Println(errFile)
			return
		}
		// 获取源文件信息内容
		sourceFileInfo, errFileInfo := sourceFile.Stat()
		if errFileInfo != nil {
			fmt.Println(errFileInfo)
			return
		}
		// 判断源文件是否为文件夹
		if sourceFileInfo.IsDir() {
			// 是文件夹, 修改路径信息
			targetFilePath := strings.Replace(sourceFilePath, sourceDirPath, targetDirPath, 1)
			// 获取文件夹对象
			_, errTargetFile := os.Open(targetFilePath)
			if errTargetFile != nil {
				// 文件夹不存在, 新建文件夹对象
				err := os.MkdirAll(targetFilePath, 0777)
				if err != nil {
					fmt.Println("文件夹已存在!!")
					return
				}
			}
		} else {
			// 不是文件夹, 复制文件信息到新文件
			// 读取原文件信息到字节切片
			fileData := make([]byte, sourceFileInfo.Size())
			_, errNum := sourceFile.Read(fileData)
			if errNum != nil {
				fmt.Println(errNum)
				return
			}
			// 关闭源文件对象
			sourceFile.Close()
			// 修改路径信息, 获取目标文件路径
			targetFilePath := strings.Replace(sourceFilePath, sourceDirPath, targetDirPath, 1)
			// 新建文件
			targetFile, errTargetFile := os.Create(targetFilePath)
			if errTargetFile != nil {
				fmt.Println(errTargetFile)
				return
			}
			// 向文件写入数据
			targetFile.Write(fileData)
			// 关闭文件对象
			targetFile.Close()
		}
	}
}
