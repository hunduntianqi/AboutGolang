package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
	使用Golang语言操作MySQL, 执行select语句查询数据
*/

// MysqlSelectConnect 定义全局变量, 存储数据库连接对象
var MysqlSelectConnect *sqlx.DB

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
	MysqlSelectConnect = databaseConn
}

// Person 定义匹配数据结构的结构体类型Person
type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"usernmae"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

func main() {
	// 定义空接口切片, 存储数据库读取的数据
	var selectData []Person
	// 使用全局数据库连接对象执行select语句
	err := MysqlSelectConnect.Select(&selectData, "select * from person")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印数据
	for _, data := range selectData {
		fmt.Println(data)
	}
	// 关闭数据库连接
	MysqlSelectConnect.Close()
}
