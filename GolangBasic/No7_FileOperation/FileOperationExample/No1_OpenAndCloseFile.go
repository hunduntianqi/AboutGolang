package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
	打开和关闭文件
		获取文件对象:
			1. func Create(path string) (file *File, err Error): 根据路径创建新的文件对象, 如果文件已存在则会清空文件数据,
																 默认权限是0666
			2. func NewFile(fd uintptr, path string) *File: 根据文件描述符创建文件对象, 文件描述符是一个非负整数, 本质上是一个索引值
			3. func Open(path string) (file *File, err Error): 根据指定路径以只读方式打开文件对象
			4. func OpenFile(path string, flag int, perm uint32) (file *File, err Error): 指定路径打开文件对象
				path ==> 目标文件路径
				flag ==> 文件读写方式
				perm ==> 操作权限
				flag可取值对应常量:
					os.O_WRONLY ==> 只写
					os.O_CREATE ==> 创建文件对象, 如果文件已存在会报错
					os.O_RDONLY ==> 只读
					os.O_RDWR ==> 可读可写
					os.O_TRUNC ==> 清空文件数据
					os.O_APPEND ==> 追加写入
			5. func (f *File) Readdir(n int) (fi []FileInfo, err error): 读取文件夹中指定 n 个文件对象的FileInfo, 返回一个切片
			6. func (f *File) Readdirnames(n int) (names []string, err error): 读取文件夹中指定 n 个文件对象的文件名, 返回一个切片
		关闭文件对象:
			func (file *File) Close() error: 关闭文件对象并返回一个错误对象
*/

func main() {
	createFile()
	newFile()
	openFile()
	openFileBy()
	readDirFile()
}

// 定义函数, 创建新文件对象 ==> func Create(path string) (file *File, err Error)
func createFile() {
	// 调用 os.Create() 函数创建文件对象
	file, err := os.Create("No7_FileOperation/FileOperationExample/Test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("创建的新文件对象内存地址为: ", &file)
	// 获取文件对象信息
	fileInfo, _ := os.Lstat("No7_FileOperation/FileOperationExample/Test.txt")
	// 新建文件对象文件名为
	fmt.Println("新建文件对象文件名为: ", fileInfo.Name())
	fmt.Println("新建文件对象是否为文件夹: ", fileInfo.Mode().IsDir())
}

// 定义函数, 根据文件描述符创建文件对象 ==> func NewFile(fd uintptr, path string) *File
func newFile() {
	file := os.NewFile(3, "No7_FileOperation/FileOperationExample/Test.txt")
	fmt.Println("打开的文件对象内存地址为: ", &file)
}

// 定义函数, 根据路径打开文件对象 ==> func Open(path string) (file *File, err Error)
func openFile() {
	file, err := os.Open("No7_FileOperation/FileOperationExample/Test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("打开的文件对象内存地址为: ", &file)
	// 关闭文件对象
	errFileClose := file.Close()
	if err != nil {
		fmt.Println(errFileClose)
		return
	}
}

// 定义函数, 以指定读写方式打开文件对象 ==> func OpenFile(path string, flag int, perm uint32) (file *File, err Error)
func openFileBy() {
	file, err := os.OpenFile("No7_FileOperation/FileOperationExample/Test.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("打开的文件对象内存地址为: ", &file)
	// 关闭文件对象
	errFileClose := file.Close()
	if err != nil {
		fmt.Println(errFileClose)
		return
	}
}

// 定义函数, 读取指定文件夹中所有文件信息
func readDirFile() {
	// 定义变量, 指定文件夹路径
	path := ""
	fmt.Println("请输入指定文件夹路径: ")
	_, err := fmt.Scan(&path)
	if err != nil {
		return
	}
	// 打开文件对象
	file, _ := os.Open(path)
	// 获取文件夹中文件个数
	fileNumber, _ := filepath.Glob(path + "//*")
	// 读取文件夹中所有文件FileInfo数据
	listFileInfo, errFileInfo := file.Readdir(len(fileNumber))
	if errFileInfo != nil {
		fmt.Println(errFileInfo)
		return
	}
	// 遍历列表
	for _, fileInfo := range listFileInfo {
		fmt.Println("fileName ==> ", fileInfo.Name())
		fmt.Println("fileSize ==> ", fileInfo.Size())
		fmt.Println("fileModTime ==> ", fileInfo.ModTime())
	}
}
