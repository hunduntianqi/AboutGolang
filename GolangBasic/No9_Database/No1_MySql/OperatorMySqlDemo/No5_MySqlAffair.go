package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
	使用Golang语言操作MySQL, 测试事务相关操作
*/

// MysqlAffairConnect 定义全局变量, 存储数据库连接对象
var MysqlAffairConnect *sqlx.DB

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
	MysqlAffairConnect = databaseConn
}

func main() {
	// 执行增删改查之前, 开启事务
	begin, err := MysqlAffairConnect.Begin()
	if err != nil {
		fmt.Println("事务开启失败!!")
		return
	}
	// 使用全局数据库连接对象执行delete语句
	result, err := MysqlAffairConnect.Exec("delete from person where user_id >=?", 4)
	if err != nil {
		fmt.Println(err)
		err := begin.Rollback() // 操作出现异常, 回滚事务
		if err != nil {
			fmt.Println("回滚事务失败")
			return
		}
		return
	}
	lineNum, errLineNum := result.RowsAffected()
	if errLineNum != nil {
		fmt.Println(errLineNum)
		err := begin.Rollback() // 操作出现异常, 回滚事务
		if err != nil {
			fmt.Println("回滚事务失败")
			return
		}
		return
	}
	fmt.Println("受影响数据为", lineNum, "行")
	// 所有操作执行完毕, 提交事务
	errCommit := begin.Commit()
	if errCommit != nil {
		fmt.Println("事务提交失败")
		return
	}
	// 关闭数据库连接
	MysqlAffairConnect.Close()
}
