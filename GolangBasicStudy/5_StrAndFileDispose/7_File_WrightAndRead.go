package main

import (
	"fmt"
	"os"
	"strings"
)

/*
	str := fmt.Sprintf(格式化字符串): 将格式化后的字符串赋值给str
*/

// 定义函数向文件写入数据
func WriteFile(path string) {
	// 打开文件
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	// 向文件写入数据
	str := fmt.Sprintf("%s", "this is one txt file")
	file.WriteString(str)
	// os.Stdout.WriteString("this is one txt file\n")
	fmt.Println(str)
	// 关闭文件
	file.Close()
}

// 定义函数读取文件数据
func ReadFile(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
	}
	// print(file.Name() + "\n")
	// fmt.Println(file.Name())
	file_stat, err_stat := os.Stat(file.Name())
	if err_stat != nil {
		fmt.Println(err_stat)
		return
	}
	// 定义字节切片
	file_data := make([]byte, file_stat.Size(), file_stat.Size()+1024) // 定义与文件字节大小一致的字节切片
	data, err_read := file.Read(file_data)                             // 返回文件数据的字节数
	if err_read != nil {
		fmt.Println("err_read=", err_read)
		return
	}
	fmt.Println("文件字节数=", data)
	fmt.Println(len(string(file_data)))
	fmt.Println("文件内容为:\n", string(file_data))
	old_filed := ""
	fmt.Println("请输入要替换的旧字段:")
	fmt.Scan(&old_filed)
	new_field := ""
	fmt.Println("请输入替换后的新字段:")
	fmt.Scan(&new_field)
	file_write, err_write := os.Create(strings.Replace(file.Name(), old_filed, new_field, 1))
	if err_write != nil {
		fmt.Println(err_write)
		return
	}
	file_write.WriteString(string(file_data))
	file_write.Close()
	file.Close()
}

// 定义函数删除文件
func RemoveFile(path string) {
	os.Remove(path)
}
func main() {
	// 定义文件路径
	str := ""
	fmt.Println("请输入文件路径:")
	fmt.Scan(&str)
	path := str
	// 调用函数写入数据
	// WriteFile(path)
	// 调用函数读取文件
	ReadFile(path)
	// 调用函数删除文件
	// RemoveFile(path)

}
