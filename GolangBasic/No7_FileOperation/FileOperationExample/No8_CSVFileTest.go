package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

/*
	CSV 文件读写:
*/

func main() {
	// 调用函数, 向 csv 文件写入数据
	//writeCSV()
	// 调用函数, 读取csv文件数据
	readCSV()
}

// 定义函数, 向CSV文件写入数据
func writeCSV() {
	// 新建一个.CSV文件
	csvFile, err := os.Create("No7_FileOperation/FileOperationExample/csvTest.csv")
	if err != nil {
		fmt.Println("文件创建失败 ==> ", err)
		return
	}
	// 创建一个 csv.Writer 对象
	write := csv.NewWriter(csvFile)
	// 创建字符串切片, 准备数据
	data := make([]string, 0)
	data = append(data, "name", "age", "sex")
	// 向文件写入数据
	errWrite := write.Write(data)
	if errWrite != nil {
		fmt.Println("写入数据失败 ==> ", err)
		return
	}
	// 定义二维字符串切片, 新增数据
	newData := make([][]string, 2)
	newData[0] = append(newData[0], "郭鹏涛", "25", "男")
	newData[1] = append(newData[1], "陈欣妮", "25", "女")
	// 将新增数据全部写入文件
	errWriteAll := write.WriteAll(newData)
	if errWriteAll != nil {
		return
	}
	// 关闭文件对象
	write.Flush()
	defer csvFile.Close()
}

// 定义函数, 从CSV文件读取数据
func readCSV() {
	// 已只读方式打开 csv 文件对象
	csvFile, errOpen := os.Open("No7_FileOperation/FileOperationExample/csvTest.csv")
	if errOpen != nil {
		fmt.Println("打开文件失败 ==> ", errOpen)
		return
	}
	// 创建 csv.Reader 对象
	read := csv.NewReader(csvFile)
	// 读取一行数据
	fmt.Println(read.Read())
	// 读取剩余所有数据
	fmt.Println(read.ReadAll())
	// 关闭文件对象
	defer csvFile.Close()
}
