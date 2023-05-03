package No2_Redis

/*
	安装第三方包:
		go get github.com/gomodule/redigo/redis
	导包:
		"github.com/gomodule/redigo/redis"
	连接Redis:
		redisConn, err := redis.Dial("tcp", "ipAddr:port", redis.DialDatabase(db: int))
			ipAddr ==> 目标主机IP地址
			port ==> redis数据库使用端口, 默认为6379
			redisConn ==> redis数据库连接对象
			redis.DialDatabase(db: int) ==> 指定要写入数据的数据库序号(0~15)
	Redis ==> string 类型数据操作:
		Redis 中 string数据用于存储一组键值对 ==> keyName Value
		设置单个键值 ==> redisConn.Do("Set", KeyName, ValueName)
		获取对应键的值 ==>  redis.ValueType(redisConn.Do("Get", KeyName)):
			==> ValueType为对应 value的类型(String / Int / Rune / Byte / Bool)
		设置多个键值 ==> redisConn.Do("MSet", KeyName1, ValueName1, KeyName2, ValueName2...)
		获取多个键的值 ==> redis.Values(redisConn.Do("MGet", KeyName1, KeyName2...)) ==> 返回数据为一个空接口切片
		设置键的过期时间 ==> redisConn.Do("expire", KeyName, timeOut) ==> timeOut单位为秒
		对应键追加值 ==> redisConn.Do("Append", KeyName, AppendValue)
	Redis ==> list 类型数据操作:
		Redis 中 list类型 用于存储多个数据, 类似一个切片
		按照指定顺序反向写入数据 到对应List ==> redisConn.Do("lpush", "KeyName", Value1, Value2, ....):
			==> 数据库中 KeyName 数据顺序 ==> ...., Value2, Value1
		按照指定顺序写入数据 到对应List ==> redisConn.Do("lpush", "KeyName", Value1, Value2, ....):
			==> 数据库中 KeyName 数据顺序 ==> Value1, Value2, ....
		在对应List中的指定数据前插入数据 ==> redisConn.Do("linsert", "KeyName", "before", oldValue, newValue)
		在对应List中的指定数据后插入数据 ==> redisConn.Do("linsert", "KeyName", "after", oldValue, newValue)
		获取对应List指定范围数据 ==> redis.Values(redisConn.Do("lrange", "KeyName", startIndex, stopIndex)):
			==> 返回数据为切片类型, 可使用 for + range 循环遍历
		修改对应List指定索引位置的数据 ==> redisConn.Do("lset", "KeyName", Index, newValue)
		删除对应List中指定的数据 ==> redisConn.Do("lrem", "KeyName", count, Value):
			count > 0:从头往尾移除
			count < 0:从尾往头移除
			count = 0:移除所有
	Redis ==> Hash 类型数据操作:
		Redis 中 Hash 类型数据用于存储一个对象, 数据结构为: objectName{field1:value1, field2:value2, ...}
		设置单个属性 ==> redisConn.Do("hset", "objectName", field, value)
		设置多个属性 ==> redisConn.Do("hmset", "objectName", field1, value1, field2, value2, ...)
		获取指定对象的所有属性名 ==> redis.Values(redisConn.Do("hkeys", "objectName"))
			==> 返回数据为切片类型, 可使用 for + range 循环遍历
		获取指定对象指定属性值 ==> redis.ValueType(redisConn.Do("hget", "objectName", field)):
			==> ValueType为对应 value的类型(String / Int / Rune / Byte / Bool)
		获取指定对象多个属性值 ==> redis.Values(redisConn.Do("hmget", "objectName", field1, field2, ...)):
			==> 返回数据为切片类型, 可使用 for + range 循环遍历
		获取指定对象所有属性值 ==> redis.Values(redisConn.Do("hvals", "objectName")):
			==> 返回数据为切片类型, 可使用 for + range 循环遍历
		删除指定对象的指定属性 ==> connRedisHash.Do("hdel", "objectName", field)
		删除指定对象及其所有属性 ==> connRedisHash.Do("del", "objectName")
	Redis ==> Set 类型数据操作:
		Redis 中 Set 类型数据为一个无序集合, 集合元素具有唯一性, 不重复, 且无修改操作
		新增 Set 集合数据 ==> connRedisSet.Do("sadd", "setName", value1, value2, ...)
		获取对应 Set 集合所有数据 ==> redis.Values(connRedisSet.Do("smembers", "setName")):
			==> 返回数据为切片类型, 可使用 for + range 循环遍历
		删除 Set 集合指定数据 ==> connRedisSet.Do("srem", "setName", value1, value2, ...):
			==> 注意: 当 Set 集合所有数据被删除时, 集合会被数据库自动删除
	Redis ==> ZSet 类型数据操作:
		Redis中 ZSet 类型数据为一个有序集合, 集合元素具有唯一性, 不重复, 且无修改操作, 每个元素都会关联一个double类型的score, 表示权重,
		通过权重将元素从小到大排序
		新增 ZSet 集合数据 ==> connRedisZSet.Do("zadd", "zsetName", score1, value1, score2, value2 ...)
		获取 ZSet 集合指定索引范围的数据 ==> redis.Values(connRedisZSet.Do("zrange", "zsetName", startIndex, stopIndex)):
			==> 返回数据为切片类型, 可使用 for + range 循环遍历
		获取 ZSet 集合指定 Score 范围的数据 ==> redis.Values(connRedisZSet.Do("zrangebyscore", "zsetName", min, max):
			==> 返回数据为切片类型, 可使用 for + range 循环遍历
		获取 ZSet 集合指定数据的Score值 ==> redis.Float64(connRedisZSet.Do("zscore", "zsetName", value))
		删除 ZSet 集合指定数据 ==> connRedisZSet.Do("zrem", "zsetName", value1, value2, ...)
		删除 ZSet 集合指定Score范围的元素 ==> connRedisZSet.Do("zremrangebyscore", "gatherName", min, max):
			==> 注意: 当 ZSet 集合所有数据被删除时, 集合会被数据库自动删除
	Redis连接池:
		创建全局Redis连接池对象 ==> var poolName *redis.Pool
		初始化连接池对象:
			poolName =&redis.Pool {
					MaxIdle: 16, ==> 设置连接池初始连接对象数量
					Dial: func() (redis.Conn, error) {
						return redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0))
					}, ==> 定义要连接的数据库
					MaxActive: 100, ==> 设置连接池最大连接对象数量, 如果为0, 表示没有限制, 按需分配
					IdleTimeout: 300 ==> 设置连接关闭时间, 单位为秒, 指定时间不使用自动关闭连接
				}
		从连接池中获取连接对象 ==> connObject := poolName.Get()
		关闭连接池 ==> poolName.Close()
*/
