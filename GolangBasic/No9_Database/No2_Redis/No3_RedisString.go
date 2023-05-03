package No2_Redis

/*
	redis数据类型: string:
        字符串类型是Redis中最为基础的数据类型, 在redis中是二进制安全的, 该类型数据可以接受
        任何格式的数据, 如图像和Json对象描述信息等, 在redis中字符串类型的Value最多可以容纳
        的数据长度为512M
        设置键值对相关命令:
            1.1 设置一组键值对:
                如果设置的键不存在则添加, 键已存在则修改对应的值:set key value
                例: set name itcast
            1.2 设置键值及过期时间(指一定时间后该键值对被清除)
                setex key seconds value (该过期时间以秒为单位)
            1.3 设置多个键值对:
                mset key1 value1 key2 value2 ...
            1.4 追加值:
                append key value:
                    向已有键key中的值追加值value
                    例:已有key:name, value:gpt
                    append name 'love cxn'
                    结果为: key:name, value:'gpt love cxn'
        获取键的值相关命令:
            1. 根据键获取值, 如果不存在则返回nil:
                get key
            2. 根据多个键获取多个值:
                mget key1 key2 ...
*/
