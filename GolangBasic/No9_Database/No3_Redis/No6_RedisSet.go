package No3_Redis

/*
	redis数据类型-set:
        无序集合
        元素为string类型
        元素具有唯一性, 不重复
        说明: 对于集合没有修改操作
        增加:
            1. 添加元素:
                sadd key member1 member2...
        获取:
            1. 返回所有元素:
                smembers key
        删除:
            1. 删除指定元素:
                srem key value1 value2 ...
*/
