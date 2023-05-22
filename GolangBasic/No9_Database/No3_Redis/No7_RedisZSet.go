package No2_Redis

/*
	redis数据类型-zset:
        sorted set, 有序集合
        元素为string类型
        元素具有唯一性, 不重复
        每个元素都会关联一个double类型的score, 表示权重, 通过权重将元素从小到大排序
        说明:没有修改操作
        增加:
            1. 添加:
                zadd key score1 member1 score2 member2 ...
        获取:
            1. 返回指定范围内的元素:
                zrange key start stop
                start, stop为元素的下标索引, 索引从左侧开始, 第一个元素为0, 索引可以是负数,
                表示从尾部开始计数, 如-1表示最后一个元素
            2. 返回score值在min和max之间的成员:
                zrangebyscore key min max
            3. 返回成员member的score值:
                zscore key member
        删除:
            1. 删除指定元素:
                zrem key member1 member2 ...
            2. 删除权重在指定范围内的元素:
                zremrangebyscore key min max
*/
