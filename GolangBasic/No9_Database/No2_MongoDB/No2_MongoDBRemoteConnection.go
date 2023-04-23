package No2_MongoDB

/*
	MongoDB远程连接设置:
        方式一: 修改mongod.cfg内容如下:
            # network interfaces
            net:
              port: 27017
              bindIpAll: true
        方式二: 修改mongod.cfg内容如下:
            # network interfaces
            net:
              port: 27017
              bindIp: 0.0.0.0
        个人电脑宿舍WiFi ipv4 地址 ==> 192.168.16.101
        个人电脑 ==> GPT-HunDunTianQi
        远程连接命令:
            mongo 远程主机ip或DNS:MongoDB端口号/数据库名 -u user -p password
*/
