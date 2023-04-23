package main

import (
	"fmt"
	"os"
)

/*
	读取文件数据:
		1. func (file *File) Read(b []byte) (n int, err Error): 读取文件数据到一个字节切片中, 当没有数据可以读取时err返回EOF
				返回值:
					n ==> 读取数据字节数
					err ==> Error对象
		2. func (file *File) ReadAt(b []byte, off int64) (n int, err Error): 从指定位置开始读取文件数据到一个字节切片中,
																				当没有数据可以读取时err返回EOF
			返回值:
				n ==> 读取数据字节数
				err ==> Error对象
	向文件写入数据:
		1. func (file *File) Write(b []byte) (n int, err Error): 将一个字节切片数据写入文件
				返回值:
					n ==> 写入文件数据的字节数
					err ==> Error对象
		2. func (file *File) WriteAt(b []byte, off int64) (n int, err Error): 从指定位置开始写入字节切片中的数据
			返回值:
				n ==> 写入文件数据的字节数
				err ==> Error对象
		3. func (file *File) WriteString(s string) (ret int, err Error): 讲一个字符串数据写入文件对象
			返回值:
				ret ==> 写入文件数据的字符数
				err ==> Error对象
*/

func main() {
	readFileData()
	writeFileData()
}

// 定义函数, 读取文件数据 ==> func (file *File) Read(b []byte) (n int, err Error)
func readFileData() {
	// 定义变量存储文件路径
	var path string
	fmt.Println("请输入目标文件路径: ")
	_, err := fmt.Scan(&path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取文件对象
	file, _ := os.Open(path)
	// 定义切片存储读取文件数据
	dataList := make([]byte, 1024)
	for {
		num, err := file.Read(dataList)
		if err != nil {
			fmt.Println("文件读取完毕: ", err)
			return
		}
		fmt.Println(string(dataList[:num]))
	}
	// 关闭文件对象
	file.Close()
}

// 定义函数, 向文件中写入数据
func writeFileData() {
	// 定义变量存储文件路径
	var path string
	fmt.Println("请输入目标文件路径: ")
	_, err := fmt.Scan(&path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 新建文件对象
	file, _ := os.OpenFile(path, os.O_WRONLY, 0666)
	// 定义切片存储文件数据
	dataList := make([]byte, 0)
	dataList = append(dataList, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	numByte, errWrite := file.Write(dataList) // 以字节切片向文件写入数据
	fmt.Println("成功写入字节数为:", numByte)
	if err != nil {
		fmt.Println(errWrite)
		return
	}

	numChar, errWriteString := file.WriteString("混沌天炁\n") // 向文件写入字符串数据
	fmt.Println("成功写入字符数为:", numChar/3)
	if err != nil {
		fmt.Println(errWriteString)
		return
	}
	// 关闭文件对象
	errCloseFile := file.Close()
	if errCloseFile != nil {
		fmt.Println(errCloseFile)
		return
	}
}
