package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
	使用Golang语言操作MySQL, 执行insert语句写入数据
*/

// MysqlInsertConnect 定义全局变量, 存储数据库连接对象
var MysqlInsertConnect *sqlx.DB

// 定义初始化函数, 初始化连接数据库
func init() {
	// 定义字符串变量, 存储连接数据库的参数信息
	connectParam := "root:2251789949gpt@tcp(127.0.0.1:3306)/golangstudy"
	// 连接数据库, 获取数据库连接对象
	databaseConn, err := sqlx.Open("mysql", connectParam)
	if err != nil {
		fmt.Println("数据库连接失败: ", err)
		return
	}
	// 给全局变量赋值
	MysqlInsertConnect = databaseConn
}

func main() {
	// 使用全局数据库连接对象执行insert语句
	exec, err := MysqlInsertConnect.Exec("insert into person(usernmae, sex, email) values (?, ?, ?)", "郭鹏涛", "男", "1729992141@qq.com")
	if err != nil {
		fmt.Println("向数据库写入数据失败: ", err)
		return
	}
	id, errId := exec.LastInsertId()
	if errId != nil {
		fmt.Println(errId)
		return
	}
	fmt.Println(id)
	// 关闭数据库连接
	MysqlInsertConnect.Close()
}
