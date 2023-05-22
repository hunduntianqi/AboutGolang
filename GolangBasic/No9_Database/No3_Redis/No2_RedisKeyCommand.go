package No2_Redis

/*
	Redis-键命令:
        Redis数据通用命令, 对Redis数据库中所有数据类型均适用
        查找键, 参数支持正则表达式:
            keys pattern
            1. 查看所有键-keys *
            2. 查看键名以某一字符开头的键-keys 字符*
        判断键是否存在, 如果存在返回1, 不存在返回0:
            exists key1
        查看键对应的value的类型(此类型为redis数据库支持的五种类型之一):
            type key
        删除键及对应的值:
            del key1 key2 ...
        设置键的过期时间, 以秒为单位:
            如果没有指定过期时间则一直存在, 直到使用del删除
            expire key seconds
        查看键的有效时间, 以秒为单位:
            ttl key
*/
