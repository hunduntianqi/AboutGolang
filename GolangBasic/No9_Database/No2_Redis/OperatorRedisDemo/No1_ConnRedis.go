package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

/*
	连接Redis
*/

func main() {
	// 创建Redis数据库连接对象
	connRedis, connErr := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0))
	if connErr != nil {
		fmt.Println("Redis数据库连接失败: ", connErr)
		return
	}
	fmt.Println("Redis数据库连接成功 ==> ", connRedis)
	// 关闭数据库连接
	connRedis.Close()
}
