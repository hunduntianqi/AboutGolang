package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

/*
	Redis 数据库连接池
*/

// 创建全局连接池对象
var poolRedis *redis.Pool

// 定义初始化函数, 初始化连接池对象
func init() {
	poolRedis = &redis.Pool{
		// 定义要连接的数据库
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0))
		},
		// 设置连接池初始连接对象数量
		MaxIdle: 16,
		// 设置连接池最大连接对象数量, 如果为0, 表示没有限制, 按需分配
		MaxActive: 100,
		// 设置连接关闭时间, 单位为秒, 指定时间不使用自动关闭连接
		IdleTimeout: 300,
	}
}

func main() {
	// 从连接池中获取连接对象
	conn := poolRedis.Get()
	// 获取键的值
	valueDatas, errGets := redis.Values(conn.Do("MGet", "Name", "Age", "Sex"))
	if errGets != nil {
		fmt.Println("获取键值失败 ==> ", errGets)
		return
	}
	fmt.Printf("Name ==> %s, Age ==> %s, Sex ==> %s\n", valueDatas[0], valueDatas[1], valueDatas[2])
	// 关闭连接对象
	conn.Close()
	// 关闭连接池
	poolRedis.Close()
}
