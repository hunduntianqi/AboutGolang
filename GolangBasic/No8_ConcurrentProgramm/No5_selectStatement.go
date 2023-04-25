package No8_ConcurrentProgramm

/*
	select语句:
		类似于switch语句, 具有一系列的case语句和一个默认的分支语句, 每个case会对应一个通道的发送或接收过程
		与switch语句相比, select语句有较多的限制, 其中最大的一条限制为: 每个case语句里必须是一个IO操作
		语法格式:
		select {
			case <- ch1:
				// ch1成功取出数据时, 要执行的语句
			case ch2 <- value:
				// 成功向ch2中存入数据时, 要执行的语句
			...
			default:
				// 以上case语句均不满足条件时执行的语句
		}
		select语句会一直等待, 直到某个case的通信操作成功完成时, 就会执行该case对应的语句, 如果一直没有case完成通信操作, select
		语句会一直等待, 如果存在default语句, 会执行default语句内容, 同时程序会从select语句后的语句中恢复
		1. select可以同时监听一个或多个channel, 直到其中一个channel ready(准备好)
		2. 如果select语句中有多个case语句同时满足条件, 则会随机选择一个case语句执行
		3. select语句可以判断channel是否存满
*/
