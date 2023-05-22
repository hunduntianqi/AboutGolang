package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

/*
	Redis List 类型数据操作
*/

func main() {
	// 连接Redis数据库
	connRedis, connErr := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(1))
	if connErr != nil {
		fmt.Println("数据库连接失败 ==> ", connErr)
		return
	}
	// 按照指定顺序写入数据到 name 中
	connRedis.Do("lpush", "name", "郭鹏涛", "陈欣妮", "GPT LOVE CXN")
	// 按照指定顺序反向写入数据到 name 中
	connRedis.Do("rpush", "name", "郭鹏涛", "陈欣妮", "GPT LOVE CXN")
	// 在指定元素前插入数据
	connRedis.Do("linsert", "name", "before", "陈欣妮", "张佳琪")
	// 在指定元素后插入数据
	connRedis.Do("linsert", "name", "after", "陈欣妮", "陶静")
	valuesSlice, errValue := redis.Values(connRedis.Do("lrange", "name", 0, 10))
	if errValue != nil {
		fmt.Println("获取数据库数据失败 ==> ", errValue)
		return
	}
	for _, listData := range valuesSlice {
		fmt.Printf("%s \n", listData)
	}
	// 修改指定索引数据
	_, errChange := connRedis.Do("lset", "name", 0, "GPT LOVE SYY")
	if errChange != nil {
		fmt.Println("修改数据失败 ==> ", errChange)
		return
	}
	// 删除列表中指定的数据
	connRedis.Do("lrem", "name", 0, "索月雅")
	// 关闭数据库连接
	connRedis.Close()
}
