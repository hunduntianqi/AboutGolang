package No2_Redis

/*
	Redis设置远程连接:
        1. 打开 redis.conf文件将 "bind 127.0.0.1" 注释掉, 或者改为"bind 0.0.0.0"
        2. 守护进程设置:
            redis.conf文件中: daemonize ==> yes: 守护进程, no: 不守护进程, 需设置为no
        3. 保护模式:
            redis.conf文件中: protected-mode ==> yes: 保护， no: 不保护, 需设置为no
        4. 设置密码(非必须):
            redis.conf文件中 ==> "#requirepass foobared" 取消注释, 改为如下数据:
            requirepass password
*/
