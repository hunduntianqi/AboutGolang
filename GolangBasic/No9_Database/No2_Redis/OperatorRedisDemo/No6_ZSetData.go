package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

/*
	Redis Hash 类型数据操作
*/
// 定义全局Redis数据库连接对象
var connRedisZSet redis.Conn

// 定义初始化函数, 初始化数据库连接
func init() {
	// 连接Redis数据库
	connRedis, connErr := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(4))
	if connErr != nil {
		fmt.Println("数据库连接失败 ==> ", connErr)
		return
	}
	connRedisZSet = connRedis
}

func main() {
	//addZSetData()
	//getZSetData()
	delZSetData()
	// 关闭数据库连接
	connRedisZSet.Close()
}

// 定义函数, 新增数据到数据库
func addZSetData() {
	_, errAdd := connRedisZSet.Do("zadd", "gatherName", 1, "One", 2, "Two", 3, "Three")
	if errAdd != nil {
		fmt.Println("数据库写入数据失败 ==> ", errAdd)
		return
	}
}

// 定义函数, 获取数据库对应集合所有数据
func getZSetData() {
	// 获取指定索引范围的数据
	valueSliceIndex, errGetValueIndex := redis.Values(connRedisZSet.Do("zrange", "gatherName", 0, 5))
	if errGetValueIndex != nil {
		fmt.Println("获取数据失败", errGetValueIndex)
		return
	}
	for _, value := range valueSliceIndex {
		fmt.Printf("%s ", value)
	}
	fmt.Println()
	fmt.Println("====================")
	// 根据Score获取指定区域的数据
	valueSliceScore, errValueScore := redis.Values(connRedisZSet.Do("zrangebyscore", "gatherName", 1, 2))
	if errValueScore != nil {
		fmt.Println("获取数据失败", errValueScore)
		return
	}
	for _, value := range valueSliceScore {
		fmt.Printf("%s ", value)
	}
	fmt.Println()
	fmt.Println("====================")
	// 获取指定数据的Score值
	valueScore, errScore := redis.Float64(connRedisZSet.Do("zscore", "gatherName", 2))
	if errScore != nil {
		fmt.Println("获取数据失败", errScore)
		return
	}
	fmt.Printf("%0.2f\n", valueScore)
}

// 定义函数, 删除数据库数据
func delZSetData() {
	// 删除 ZSet 集合指定数据
	_, errDel := connRedisZSet.Do("zrem", "gatherName", 2, 3)
	if errDel != nil {
		fmt.Println("数据库删除数据失败 ==> ", errDel)
		return
	}
	// 删除 ZSet 集合指定Score范围的元素
	_, errDelScore := connRedisZSet.Do("zremrangebyscore", "gatherName", 1, 5)
	if errDelScore != nil {
		fmt.Println("数据库删除数据失败 ==> ", errDelScore)
		return
	}
}
