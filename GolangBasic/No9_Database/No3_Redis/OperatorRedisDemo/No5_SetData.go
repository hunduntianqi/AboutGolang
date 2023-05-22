package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

/*
	Redis Hash 类型数据操作
*/
// 定义全局Redis数据库连接对象
var connRedisSet redis.Conn

// 定义初始化函数, 初始化数据库连接
func init() {
	// 连接Redis数据库
	connRedis, connErr := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(3))
	if connErr != nil {
		fmt.Println("数据库连接失败 ==> ", connErr)
		return
	}
	connRedisSet = connRedis
}

func main() {
	//addSetData()
	//getSetData()
	delSetData()
	// 关闭数据库连接
	connRedisSet.Close()
}

// 定义函数, 新增数据到数据库
func addSetData() {
	_, errAdd := connRedisSet.Do("sadd", "gatherName", 1, 2, 3, 3, 5, 5)
	if errAdd != nil {
		fmt.Println("数据库写入数据失败 ==> ", errAdd)
		return
	}
}

// 定义函数, 获取数据库对应集合所有数据
func getSetData() {
	valueSlice, errGetValue := redis.Values(connRedisSet.Do("smembers", "gatherName"))
	if errGetValue != nil {
		fmt.Println("获取数据失败", errGetValue)
		return
	}
	for _, value := range valueSlice {
		fmt.Printf("%s ", value)
	}
	fmt.Println()
}

// 定义函数, 删除数据库数据
func delSetData() {
	numList := make([]int, 0)
	numList = append(numList, 1, 2, 3, 3, 5, 5)
	_, errDel := connRedisSet.Do("srem", "gatherName", 1, 2, 3, 5, numList)
	if errDel != nil {
		fmt.Println("数据库删除数据失败 ==> ", errDel)
		return
	}
}
