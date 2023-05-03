package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

/*
	Redis String类型数据操作
*/

func main() {
	// 创建 Redis 数据库连接对象
	connRedis, connErr := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0))
	if connErr != nil {
		fmt.Println("Redis数据库连接失败 ==> ", connErr)
		return
	}
	// 设置单个键值
	_, errSet := connRedis.Do("Set", "Name", "郭鹏涛")
	if errSet != nil {
		fmt.Println("设置键值失败 ==> ", errSet)
		return
	}
	// 获取键的值
	valueData, errGet := redis.String(connRedis.Do("Get", "Name"))
	if errGet != nil {
		fmt.Println("获取键值失败 ==> ", errGet)
		return
	}
	fmt.Println("Name ==> ", valueData)

	// 设置多个键值
	_, errSets := connRedis.Do("MSet", "Name", "郭鹏涛", "Age", 25, "Sex", "男")
	if errSets != nil {
		fmt.Println("设置键值失败 ==> ", errSet)
		return
	}
	connRedis.Do("Append", "Name", "Love陈欣妮")
	// 获取多个键的值
	// 获取键的值
	valueDatas, errGets := redis.Values(connRedis.Do("MGet", "Name", "Age", "Sex"))
	if errGets != nil {
		fmt.Println("获取键值失败 ==> ", errGets)
		return
	}
	fmt.Printf("Name ==> %s, Age ==> %s, Sex ==> %s\n", valueDatas[0], valueDatas[1], valueDatas[2])
	connRedis.Flush()
	// 关闭连接
	connRedis.Close()
}
