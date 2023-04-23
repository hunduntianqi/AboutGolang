package No3_Redis

/*
	redis数据类型-list:
        列表的元素类型为string, 元素按照插入顺序排序
        增加:
            1. 在左侧插入数据:
                lpush key value1 value2 ...
            2. 在右侧插入数据:
                rpush key value1 value2 ...
            3. 在指定元素的前或后插入新元素
                linsert key before或after 现有元素 新元素
        获取:
            1. 返回列表里指定范围内的元素:
                lrange key start stop
                start, stop为元素的下标索引,索引从左侧开始, 第一个元素为0,
                索引可以是负数, 表示从尾部开始计数, -1表示最后一个元素
            2. 设置指定索引位置的元素值:
                lset key index value
        删除:
            1. 删除指定元素:
                将列表中前count次出现的元素值为value的元素溢出==移除
                lrem key count value
                count > 0:从头往尾移除
                count < 0:从尾往头移除
                count = 0:移除所有
*/
