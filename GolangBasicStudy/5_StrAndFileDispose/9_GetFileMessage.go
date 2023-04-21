package main

import (
	"fmt"
	"os"
)

/*
	通过os.Stat方法，我们可以获取文件的信息，比如文件大小、名字等
		func main() {
			fi,err:=os.Stat(文件路径)
			if err ==nil {
				fmt.Println("name:",fi.Name()) // 获取文件名称
				fmt.Println("size:",fi.Size()) // 获取文件大小, 以字节形式
				fmt.Println("is dir:",fi.IsDir()) // 判断文件是否为一个文件夹, 是则返回true, 不是则返回false
				fmt.Println("mode::",fi.Mode()) // 获取文件读写模式
				fmt.Println("modTime:",fi.ModTime()) // 获取文件创建时间
			}
		}
*/

func main() {
	// 打开文件
	file, err := os.Stat("D:\\Users\\Administrator\\Desktop\\小练习\\test\\蜘蛛侠.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件名: ", file.Name())
	fmt.Printf("文件大小: %d字节\n", file.Size())
	fmt.Println("该文件是否为文件夹: ", file.IsDir())
	fmt.Println(file.Mode())
	fmt.Println("文件创建时间为: ", file.ModTime())
	fmt.Println(file.Sys())

}
