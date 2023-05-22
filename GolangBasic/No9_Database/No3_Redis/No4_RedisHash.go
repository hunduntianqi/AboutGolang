package No2_Redis

/*
	redis数据类型-hash:
        hash用于存储对象, 对象的结构为属性, 值
        值的类型为string
        增加:
            1. 设置单个属性:
                hset key field value
                例:hset user name itheima
            2. 设置多个属性:
                hmset key field1 value1 field2 value2
                例:hmset user name gpt age 24 sex man
        获取:
            1. 获取指定键所有的属性:
                hkeys key
            2. 获取指定键一个属性的值:
                hget key field
            3. 获取指定键多个属性的值:
                hmget key field1 field2...
            4. 获取所有属性的值:
                hvals key
        删除:
            1. 删除整个hash键及值, 使用del命令
            2. 删除属性, 属性对应的值会被一起删除:
                hdel key field1 field2 ...
*/
