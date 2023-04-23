package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
	使用Golang语言操作MySQL, 执行update语句修改
*/

// MysqlUpdateConnect 定义全局变量, 存储数据库连接对象
var MysqlUpdateConnect *sqlx.DB

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
	MysqlUpdateConnect = databaseConn
}

func main() {
	// 使用全局数据库连接对象执行update语句
	result, err := MysqlUpdateConnect.Exec("update person set user_id=4 where user_id=?", 9)
	if err != nil {
		fmt.Println(err)
		return
	}
	lineNum, errLineNum := result.RowsAffected()
	if errLineNum != nil {
		fmt.Println(errLineNum)
		return
	}
	fmt.Println("受影响数据为", lineNum, "行")
	// 关闭数据库连接
	MysqlUpdateConnect.Close()
}
