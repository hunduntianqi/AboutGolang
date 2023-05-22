package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

/*
	Redis Hash 类型数据操作
*/
// 定义全局Redis数据库连接对象
var connRedisHash redis.Conn

// 定义初始化函数, 初始化数据库连接
func init() {
	// 连接Redis数据库
	connRedis, connErr := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(2))
	if connErr != nil {
		fmt.Println("数据库连接失败 ==> ", connErr)
		return
	}
	connRedisHash = connRedis
}

func main() {
	//addHashData()
	delHashData()
	// 关闭数据库连接
	connRedisHash.Close()
}

// 定义函数, 新增数据到数据库
func addHashData() {
	// 创建一个对象, 并定义单个属性和属性值
	connRedisHash.Do("hset", "User", "name", "郭鹏涛")
	// 定义多组属性
	connRedisHash.Do("hmset", "User", "age", 25, "sex", "男")
}

// 定义函数, 查询数据库数据
func getHashData() {
	// 获取指定对象的所有属性名
	keyList, keyErr := redis.Values(connRedisHash.Do("hkeys", "User"))
	if keyErr != nil {
		fmt.Println("获取对象属性失败 ==> ", keyErr)
	}
	for _, keyName := range keyList {
		fmt.Printf("%s ", keyName)
	}
	fmt.Println()
	// 获取指定对象指定属性值
	value1, _ := redis.String(connRedisHash.Do("hget", "User", "name"))
	fmt.Println("name ==> ", value1)
	// 获取指定对象多个属性值
	valueList, _ := redis.Values(connRedisHash.Do("hmget", "User", "name", "age"))
	for _, value := range valueList {
		fmt.Printf("%s\n", value)
	}
	// 获取指定对象所有属性值
	valueListAll, _ := redis.Values(connRedisHash.Do("hvals", "User"))
	for _, value := range valueListAll {
		fmt.Printf("%s\n", value)
	}
}

// 定义函数, 删除数据库数据
func delHashData() {
	// 删除指定对象的指定属性
	connRedisHash.Do("hdel", "User", 1)
	// 删除指定对象及其所有属性
	connRedisHash.Do("del", "User")
}
