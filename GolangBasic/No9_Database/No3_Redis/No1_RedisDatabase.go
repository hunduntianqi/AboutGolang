package No2_Redis

/*
	Redis是一个高性能的key-value数据库, 属于NoSQL数据库, 不支持sql语法, 存储数据为kv形式
    Redis数据库特点:
        1. 支持数据持久化, 可以将内存中的数据保存在磁盘中, 重启时可以再次加载进行使用
        2. Redis数据库不仅支持key-value类型数据, 还提供list, set, zset, hash等数据结构的存储
        3. Redis支持数据备份, 即master-slave模式的数据备份
    Redis数据库优势:
        1. 性能极高-Redis能读的速度是110000次/s, 写的速度是81000次/s
        2. 丰富的数据类型-Redis支持二进制案例的Strings, Lists, Hashes, Sets及Ordered Sets数据类型操作
        3. 原子特性-Redis所有的操作都是原子性的, 同时Redis还支持对几个操作合并后的原子性执行
        4. 丰富的特性-Redis还支持publish/subscribe, 通知, key过期等特性
    Redis应用场景:
        1. 用来做缓存(ehcache/memcached)-redis所有数据是放在内存中的(内存数据库)
        2. 可以在某些特定场景下替代传统数据库-比如社交类的应用
        3. 在一些大型系统中, 巧妙地实现一些特定的功能:session共享, 购物车等
    Redis配置文件-redis.conf(windows为redis.windows.conf):
        核心配置选项:
            1. 绑定IP:如果需要远程访问, 可将此行注释或绑定一个真实IP
                bind:127.0.0.1
            2. 端口:默认为6379
                port:6379
            3. 是否以守护进程进行:
                daemonize yes
                3.1 如果以守护进程运行, 则不会在命令行阻塞, 类似于服务
                3.2 如果以非守护进程运行则当前终端被阻塞
                3.3 设置为yes表示守护进程, 设置为no表示非守护进程
            4. 数据文件:
                dbfilename dump.rdb
            5. 数据文件存储路径:
                dir 文件存储路径
            6. 日志文件:
                logfile/var/log/redis/redis-server.log
            7. 数据库:默认有16个
                database:16
            8. 主从复制, 类似于双机备份:
                slaveof
    启动服务器端和客户端:
        服务器端启动命令:telnet Ip 端口
        客户端启动:cmd命令窗口, redis-cli
        清空redis数据库:flushall
*/
